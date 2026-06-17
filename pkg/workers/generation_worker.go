package workers

import (
	"context"
	"trainers-manager/pkg/logger"

	"trainers-manager/internal/repo"
	"trainers-manager/internal/repo/workers"

	"github.com/google/uuid"
)

type GenEvent struct {
	TaskID uuid.UUID
	Status string
	Error  *string
}

type GenWorker struct {
	Event chan GenEvent
	l     *logger.Logger
	r     repo.GenerationTaskRepo
}

func NewGenWorker(l *logger.Logger, eventCh chan GenEvent, r repo.GenerationTaskRepo) *GenWorker {
	return &GenWorker{
		Event: eventCh,
		l:     l,
		r:     r,
	}
}

func (w *GenWorker) Run(ctx context.Context) {
	for {
		select {
		case event := <-w.Event:
			w.handle(ctx, event)

		case <-ctx.Done():
			w.l.Info("worker stopped")
			return
		}
	}
}

func (w *GenWorker) handle(ctx context.Context, e GenEvent) {
	switch e.Status {
	case "CREATED":
		_ = w.CreateGenerationTask(ctx, e)

	case "SENT", "PROCESSING", "ERROR", "DONE":
		_ = w.UpdateGenerationTask(ctx, e)
	}
}

func (w *GenWorker) CreateGenerationTask(ctx context.Context, e GenEvent) error {
	return w.r.CreateGenerationTask(ctx, workers.GenerationTask{ID: e.TaskID, Status: e.Status})
}

func (w *GenWorker) UpdateGenerationTask(ctx context.Context, e GenEvent) error {
	return w.r.UpdateGenerationTask(ctx, workers.UpdateGenerationTask{ID: e.TaskID, Status: &e.Status, Error: e.Error})
}

// у меня будет воркер который будет получать ивент и делать запись в бд о генерации а так же запускать воркера для обработки этих же записей далее
