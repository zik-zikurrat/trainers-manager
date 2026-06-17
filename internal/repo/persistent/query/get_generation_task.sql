SELECT id, status, error, created_at, updated_at FROM generation_tasks
WHERE id = $1;
