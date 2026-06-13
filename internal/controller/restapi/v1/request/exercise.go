package request

type Exercise struct {
	Muscle      string `json:"muscle" validate:"required"`
	Position    string `json:"position" validate:"required"`
	Description string `json:"description"  validate:"required"`
}

type UpdateExercise struct {
	Muscle      *string `json:"muscle"`
	Position    *string `json:"position"`
	Description *string `json:"description"`
}
