UPDATE training_group
SET name = $1, accent_cycle = $2, skill_cycle = $3
WHERE id = $4;
