ALTER TABLE training_plan
    ALTER COLUMN skills TYPE TEXT USING array_to_string(skills, ', ');
