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

	"github.com/google/uuid"
)

type Generator struct {
	cfg    *config.Config
	client *http.Client
}

func New(cfg *config.Config) *Generator {
	return &Generator{
		cfg:    cfg,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

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

func (g *Generator) Generate(ctx context.Context, prompt usecase.GeneratePrompt) (usecase.GeneratedPlan, error) {
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

	resp, err := g.client.Do(httpReq)
	if err != nil {
		return usecase.GeneratedPlan{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return usecase.GeneratedPlan{}, fmt.Errorf("read body: %w", err)
	}
	var chatResp chatResponse
	if err := json.Unmarshal(bodyBytes, &chatResp); err != nil {
		return usecase.GeneratedPlan{}, fmt.Errorf("decode response: %w (raw: %s)", err, string(bodyBytes))
	}

	if resp.StatusCode != http.StatusOK {
		msg := "unknown"
		if chatResp.Error != nil {
			msg = chatResp.Error.Message
		}
		return usecase.GeneratedPlan{}, fmt.Errorf("llm api error (%d): %s", resp.StatusCode, msg)
	}
	if len(chatResp.Choices) == 0 {
		return usecase.GeneratedPlan{}, fmt.Errorf("llm returned no choices")
	}

	content := chatResp.Choices[0].Message.Content

	content = stripCodeFences(content)

	var parsed llmPlan
	if err := json.Unmarshal([]byte(content), &parsed); err != nil {
		return usecase.GeneratedPlan{}, fmt.Errorf("parse llm json: %w (raw: %s)", err, content)
	}

	ids := make([]uuid.UUID, 0, len(parsed.ExerciseIDs))
	for _, s := range parsed.ExerciseIDs {
		id, err := uuid.Parse(strings.TrimSpace(s))
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}

	return usecase.GeneratedPlan{
		ExerciseIDs: ids,
		PlanText:    parsed.Plan,
	}, nil
}

func stripCodeFences(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "```json")
	s = strings.TrimPrefix(s, "```")
	s = strings.TrimSuffix(s, "```")
	return strings.TrimSpace(s)
}
