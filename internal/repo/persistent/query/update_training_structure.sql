UPDATE training_structure
SET structure = $1, updated_at = NOW()
WHERE id = $2;
