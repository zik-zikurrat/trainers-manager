INSERT INTO training_plan (plan, status, train_id, group_id, accent, skills, training_structure_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at;
