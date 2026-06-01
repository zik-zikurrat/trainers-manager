INSERT INTO exercises (muscle_group, description)
VALUES ($1, $2)
RETURNING id;
