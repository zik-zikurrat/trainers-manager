package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type HistoryAction string

const (
	HistoryActionCreate HistoryAction = "CREATE"
	HistoryActionUpdate HistoryAction = "UPDATE"
	HistoryActionDelete HistoryAction = "DELETE"
)

type TrainingPlanHistory struct {
	ID        uuid.UUID
	PlanID    uuid.UUID
	Action    HistoryAction
	Snapshot  json.RawMessage
	CreatedAt time.Time
}
