INSERT INTO training_group (name, accent_cycle, skill_cycle)
VALUES ($1, $2, $3)
RETURNING id;
