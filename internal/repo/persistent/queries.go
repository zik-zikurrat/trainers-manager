package persistent

import _ "embed"

//go:embed query/insert_training_structure.sql
var insertTrainingStructureQuery string
