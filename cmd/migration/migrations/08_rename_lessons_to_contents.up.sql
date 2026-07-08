ALTER TABLE lessons RENAME TO contents;

ALTER TABLE contents
    ADD COLUMN IF NOT EXISTS content_type VARCHAR(30) NOT NULL DEFAULT 'video';

-- Indexing speeds up ?type= filtering and title searching
CREATE INDEX IF NOT EXISTS idx_contents_content_type ON contents (content_type);
CREATE INDEX IF NOT EXISTS idx_contents_title         ON contents USING gin(to_tsvector('simple', title));

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_lessons_category') THEN
        ALTER INDEX idx_lessons_category RENAME TO idx_contents_category;
    END IF;
    IF EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_lessons_level') THEN
        ALTER INDEX idx_lessons_level RENAME TO idx_contents_level;
    END IF;
    IF EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_lessons_is_free') THEN
        ALTER INDEX idx_lessons_is_free RENAME TO idx_contents_is_free;
    END IF;
END $$;