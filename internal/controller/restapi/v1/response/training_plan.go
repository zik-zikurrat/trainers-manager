package response

import (
	"time"
	"trainers-manager/internal/entity"
)

type TrainingPlan struct {
	Plan      string    `json:"plan"`
	Status    string    `json:"status"`
	Accent    string    `json:"accent"`
	Skills    string    `json:"skills"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToTrainingPlan(trainingPlan entity.TrainingPlan) TrainingPlan {
	return TrainingPlan{
		Plan:      trainingPlan.Plan,
		Status:    trainingPlan.Status,
		Accent:    trainingPlan.Accent,
		Skills:    trainingPlan.Skills,
		CreatedAt: trainingPlan.CreatedAt,
		UpdatedAt: trainingPlan.UpdatedAt,
	}
}
