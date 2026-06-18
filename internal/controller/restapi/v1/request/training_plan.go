package request

type UpdateTrainingPlan struct {
	Plan   *string `json:"plan"`
	Accent *string `json:"accent"`
	Skills *string `json:"skills"`
}
