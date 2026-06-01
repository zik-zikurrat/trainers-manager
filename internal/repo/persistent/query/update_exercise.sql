UPDATE exercises
SET muscle_group = $1,
    description  = $2,
    updated_at   = NOW()
WHERE id = $3;
