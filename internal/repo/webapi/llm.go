package webapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"trainers-manager/internal/config"
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/workers"

	"github.com/google/uuid"
)

type Generator struct {
	cfg        *config.Config
	client     *http.Client
	genEventCh chan workers.GenEvent
}

func New(cfg *config.Config, genEventCh chan workers.GenEvent) *Generator {
	return &Generator{
		cfg:        cfg,
		client:     &http.Client{Timeout: 60 * time.Second},
		genEventCh: genEventCh,
	}
}

var _ usecase.PlanGenerator = (*Generator)(nil)

type chatRequest struct {
	Model          string        `json:"model"`
	Messages       []chatMessage `json:"messages"`
	ResponseFormat *respFormat   `json:"response_format,omitempty"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type respFormat struct {
	Type string `json:"type"`
}

type chatResponse struct {
	Choices []struct {
		Message chatMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type llmPlan struct {
	ExerciseIDs []string `json:"exercise_ids"`
	Plan        string   `json:"plan"`
}

func (g *Generator) Generate(ctx context.Context, prompt usecase.GeneratePrompt, taskID uuid.UUID) (usecase.GeneratedPlan, error) {
	reqBody := chatRequest{
		Model: g.cfg.LLM.Model,
		Messages: []chatMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: buildUserPrompt(prompt)},
		},
		ResponseFormat: &respFormat{Type: "json_object"},
	}

	raw, err := json.Marshal(reqBody)
	if err != nil {
		return usecase.GeneratedPlan{}, fmt.Errorf("marshal request: %w", err)
	}

	url := strings.TrimRight(g.cfg.LLM.BaseURL, "/") + "/chat/completions"
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(raw))
	if err != nil {
		return usecase.GeneratedPlan{}, fmt.Errorf("new request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+g.cfg.LLM.APIKey)

	g.genEventCh <- workers.GenEvent{
		TaskID: taskID,
		Status: "SENT",
		Error:  nil,
	}
	resp, err := g.client.Do(httpReq)
	if err != nil {
		errMsg := err.Error()
		g.genEventCh <- workers.GenEvent{
			TaskID: taskID,
			Status: "ERROR",
			Error:  &errMsg,
		}
		return usecase.GeneratedPlan{}, fmt.Errorf("do request: %w", err)
	}
	g.genEventCh <- workers.GenEvent{
		TaskID: taskID,
		Status: "PROCESSING",
		Error:  nil,
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := err.Error()
		g.genEventCh <- workers.GenEvent{
			TaskID: taskID,
			Status: "ERROR",
			Error:  &errMsg,
		}
		return usecase.GeneratedPlan{}, fmt.Errorf("read body: %w", err)
	}
	var chatResp chatResponse
	if err := json.Unmarshal(bodyBytes, &chatResp); err != nil {
		errMsg := err.Error()
		g.genEventCh <- workers.GenEvent{
			TaskID: taskID,
			Status: "ERROR",
			Error:  &errMsg,
		}
		return usecase.GeneratedPlan{}, fmt.Errorf("decode response: %w (raw: %s)", err, string(bodyBytes))
	}

	if resp.StatusCode != http.StatusOK {
		msg := "unknown"
		if chatResp.Error != nil {
			msg = chatResp.Error.Message
		}
		err = fmt.Errorf("llm api error (%d): %s", resp.StatusCode, msg)
		errMsg := err.Error()
		g.genEventCh <- workers.GenEvent{
			TaskID: taskID,
			Status: "ERROR",
			Error:  &errMsg,
		}
		return usecase.GeneratedPlan{}, err
	}
	if len(chatResp.Choices) == 0 {
		err = fmt.Errorf("llm returned no choices")

		errMsg := err.Error()
		g.genEventCh <- workers.GenEvent{
			TaskID: taskID,
			Status: "ERROR",
			Error:  &errMsg,
		}
		return usecase.GeneratedPlan{}, err
	}

	content := chatResp.Choices[0].Message.Content

	content = stripCodeFences(content)

	var parsed llmPlan
	if err := json.Unmarshal([]byte(content), &parsed); err != nil {
		err = fmt.Errorf("parse llm json: %w (raw: %s)", err, content)
		errMsg := err.Error()
		g.genEventCh <- workers.GenEvent{
			TaskID: taskID,
			Status: "ERROR",
			Error:  &errMsg,
		}
		return usecase.GeneratedPlan{}, err
	}

	ids := make([]uuid.UUID, 0, len(parsed.ExerciseIDs))
	for _, s := range parsed.ExerciseIDs {
		id, err := uuid.Parse(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	g.genEventCh <- workers.GenEvent{
		TaskID: taskID,
		Status: "DONE",
		Error:  nil,
	}
	return usecase.GeneratedPlan{
		ExerciseIDs: ids,
		PlanText:    cleanPlanText(parsed.Plan),
	}, nil
}

func stripCodeFences(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "```json")
	s = strings.TrimPrefix(s, "```")
	s = strings.TrimSuffix(s, "```")
	return strings.TrimSpace(s)
}
