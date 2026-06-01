CREATE TABLE IF NOT EXISTS training_structure (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    structure TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS exercises (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    muscle VARCHAR(50), 
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS training (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS training_group (    
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name         VARCHAR(50) UNIQUE NOT NULL,
    accent_cycle TEXT[] NOT NULL,
    skill_cycle  TEXT[] NOT NULL
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS training_plan (
    id                    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    plan                  TEXT,
    train_id              UUID NOT NULL REFERENCES training(id),
    group_id              UUID NOT NULL REFERENCES training_group(id),  
    training_structure_id UUID NOT NULL REFERENCES training_structure(id),
    accent                VARCHAR(50),
    skills                TEXT[],
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS training_exercises (
    training_id UUID REFERENCES training(id),
    exercise_id UUID REFERENCES exercises(id),
    PRIMARY KEY (training_id, exercise_id)
);

CREATE TABLE training_plan_history (
    id         UUID NOT NULL DEFAULT gen_random_uuid(),
    plan_id    UUID NOT NULL REFERENCES training_plan(id), 
    action     TEXT NOT NULL,
    snapshot   JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
) PARTITION BY RANGE (created_at);

CREATE OR REPLACE FUNCTION ensure_history_partitions(ahead int DEFAULT 1)
RETURNS void
LANGUAGE plpgsql AS $$
DECLARE
    base_start date := date_trunc('quarter', now())::date; 
    i          int;
    p_start    date;
    p_end      date;
    p_name     text;
BEGIN
    FOR i IN 0..ahead LOOP
        p_start := base_start + make_interval(months => i * 3);
        p_end   := p_start    + make_interval(months => 3);
        p_name  := format('training_plan_history_%sq%s',
                          to_char(p_start, 'YYYY'),
                          to_char(p_start, 'Q'));

        EXECUTE format(
            'CREATE TABLE IF NOT EXISTS %I PARTITION OF training_plan_history '
            'FOR VALUES FROM (%L) TO (%L)',
            p_name, p_start, p_end
        );
    END LOOP;
END;
$$;

SELECT ensure_history_partitions(1);
