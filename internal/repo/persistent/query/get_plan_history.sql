SELECT id, plan_id, action, snapshot, created_at
FROM training_plan_history
WHERE plan_id = $1
ORDER BY created_at DESC;
