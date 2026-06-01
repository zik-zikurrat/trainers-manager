package persistent

import _ "embed"

//go:embed query/insert_training_structure.sql
var insertTrainingStructureQuery string

//go:embed query/update_training_structure.sql
var updateTrainingStructureQuery string

//go:embed query/delete_training_structure.sql
var deleteTrainingStructureQuery string

//go:embed query/insert_exercise.sql
var insertExerciseQuery string

//go:embed query/update_exercise.sql
var updateExerciseQuery string

//go:embed query/delete_exercise.sql
var deleteExerciseQuery string

//go:embed query/list_exercises.sql
var listExercisesQuery string

//go:embed query/insert_training.sql
var insertTrainingQuery string

//go:embed query/link_exercises.sql
var linkExercisesQuery string

//go:embed query/insert_training_plan.sql
var insertTrainingPlanQuery string

//go:embed query/update_training_plan.sql
var updateTrainingPlanQuery string

//go:embed query/insert_plan_history.sql
var insertPlanHistoryQuery string

//go:embed query/get_training_plan.sql
var getTrainingPlanQuery string

//go:embed query/get_plan_history.sql
var getPlanHistoryQuery string

//go:embed query/recent_plans.sql
var getRecentPlans string

//go:embed query/get_structure.sql
var getTrainingStructure string
