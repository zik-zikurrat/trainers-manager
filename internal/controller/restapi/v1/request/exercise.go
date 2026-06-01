package request

type Exercise struct {
	MuscleGroup []string `json:"muscle_group" validate:"required,min=1,dive,required"`
	Description string   `json:"description"  validate:"required"`
}
