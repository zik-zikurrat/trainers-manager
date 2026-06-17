UPDATE generation_tasks
SET status = $1, error = $2, updated_at = NOW()
WHERE id = $3;
