INSERT INTO training_exercises (training_id, exercise_id)
SELECT $1, unnest($2::uuid[])
ON CONFLICT DO NOTHING;
