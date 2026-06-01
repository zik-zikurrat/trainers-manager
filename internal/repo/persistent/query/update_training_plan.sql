UPDATE training_plan
SET plan       = $1,
    accent     = $2,
    skills     = $3,
    updated_at = NOW()
WHERE id = $4
RETURNING train_id, training_structure_id, created_at, updated_at;
