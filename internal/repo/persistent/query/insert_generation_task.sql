INSERT INTO generation_tasks (id, status, error, created_at)
VALUES ($1, $2, $3, NOW())
RETURNING id;
