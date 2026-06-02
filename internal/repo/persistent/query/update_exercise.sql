UPDATE exercises
SET muscle = $1,
    description  = $2,
    updated_at   = NOW()
WHERE id = $3;
