package request

type TrainingGroup struct {
	Name        string   `json:"name" validate:"required,max=50"`
	AccentCycle []string `json:"accent_cycle" validate:"required,min=1,unique,dive,required"`
	SkillCycle  []string `json:"skill_cycle"  validate:"required,min=1,unique,dive,required"`
}
