UPDATE exercises
SET muscle = $1,
    description  = $2,
    position = $3,
    updated_at   = NOW()
WHERE id = $4;
