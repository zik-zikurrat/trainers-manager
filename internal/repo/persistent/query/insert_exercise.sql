INSERT INTO exercises (muscle, description)
VALUES ($1, $2)
RETURNING id;
