SELECT id, name, accent_cycle, skill_cycle
FROM training_group
WHERE name = $1;
