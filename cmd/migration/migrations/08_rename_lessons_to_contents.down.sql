DROP INDEX IF EXISTS idx_contents_title;
DROP INDEX IF EXISTS idx_contents_content_type;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_contents_category') THEN
        ALTER INDEX idx_contents_category RENAME TO idx_lessons_category;
    END IF;
    IF EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_contents_level') THEN
        ALTER INDEX idx_contents_level RENAME TO idx_lessons_level;
    END IF;
    IF EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_contents_is_free') THEN
        ALTER INDEX idx_contents_is_free RENAME TO idx_lessons_is_free;
    END IF;
END $$;

ALTER TABLE contents DROP COLUMN IF EXISTS content_type;
ALTER TABLE contents RENAME TO lessons;