INSERT INTO training_plan (plan, train_id, accent, skills, training_structure_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at;
