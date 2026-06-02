ALTER TABLE training_plan
    ALTER COLUMN skills TYPE TEXT[] USING string_to_array(skills, ', ');
