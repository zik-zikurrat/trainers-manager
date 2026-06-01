SELECT id, plan, train_id, accent, skills, training_structure_id, created_at, updated_at
FROM training_plan
WHERE id = $1;
