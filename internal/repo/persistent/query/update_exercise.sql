UPDATE exercises
SET
    muscle = COALESCE($1, muscle),
    position = COALESCE($2, position),
    description = COALESCE($3, description),
    updated_at   = NOW()
WHERE id = $4;
