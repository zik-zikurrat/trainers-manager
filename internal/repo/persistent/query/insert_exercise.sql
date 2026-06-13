INSERT INTO exercises (muscle, description, position)
VALUES ($1, $2, $3)
RETURNING id;
