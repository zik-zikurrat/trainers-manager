CREATE TABLE IF NOT EXISTS training_plan (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    plan       TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS training_structure (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    structure  TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS training_plan_history (
    id         UUID NOT NULL DEFAULT gen_random_uuid(),
    plan_id    UUID NOT NULL REFERENCES training_plan(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id, created_at)
) PARTITION BY RANGE (created_at);

CREATE TABLE training_plan_history_2026_01
    PARTITION OF training_plan_history
    FOR VALUES FROM ('2026-01-01') TO ('2026-02-01');

CREATE TABLE IF NOT EXISTS training (
    id                    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    plan_id               UUID NOT NULL REFERENCES training_plan(id),
    training_structure_id UUID NOT NULL REFERENCES training_structure(id),
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_training_plan_id              ON training(plan_id);
CREATE INDEX idx_training_structure_id         ON training(training_structure_id);
CREATE INDEX idx_training_plan_history_plan_id ON training_plan_history(plan_id, created_at DESC);
