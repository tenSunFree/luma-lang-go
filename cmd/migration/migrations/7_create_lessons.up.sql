CREATE TABLE IF NOT EXISTS lessons (
    -- Basic Information
    id               VARCHAR(100)  PRIMARY KEY,
    title            VARCHAR(255)  NOT NULL,
    subtitle         VARCHAR(255)  NOT NULL DEFAULT '',
    description      TEXT          NOT NULL DEFAULT '',
    cover_url        TEXT          NOT NULL DEFAULT '',
    duration_ms      INTEGER       NOT NULL DEFAULT 0,
    level            VARCHAR(50)   NOT NULL DEFAULT '',
    category         VARCHAR(100)  NOT NULL DEFAULT '',
    tags             TEXT[]        NOT NULL DEFAULT '{}',
    is_free          BOOLEAN       NOT NULL DEFAULT true,
    view_count       INTEGER       NOT NULL DEFAULT 0,
    captions_version INTEGER       NOT NULL DEFAULT 1,
    -- JSONB fields (best performance for reading entire records; MVP does not require cross-table JOINs)
    playback         JSONB         NOT NULL DEFAULT '{}'::jsonb,
    captions         JSONB         NOT NULL DEFAULT '[]'::jsonb,
    vocabulary_items JSONB         NOT NULL DEFAULT '[]'::jsonb,
    -- Timestamp
    created_at       TIMESTAMPTZ   NOT NULL,
    updated_at       TIMESTAMPTZ   NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_lessons_category ON lessons (category);
CREATE INDEX IF NOT EXISTS idx_lessons_level    ON lessons (level);
CREATE INDEX IF NOT EXISTS idx_lessons_is_free  ON lessons (is_free);