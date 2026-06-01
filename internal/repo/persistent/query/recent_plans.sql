SELECT accent, skills, created_at
FROM training_plan
WHERE training_structure_id = $1
ORDER BY created_at DESC
LIMIT $2;
