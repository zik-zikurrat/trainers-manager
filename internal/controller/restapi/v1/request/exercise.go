package request

type Exercise struct {
	MuscleGroup string `json:"muscle_group" validate:"required"`
	Description string `json:"description"  validate:"required"`
}
