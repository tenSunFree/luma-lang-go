CREATE UNIQUE INDEX IF NOT EXISTS idx_live_courses_one_active_live_per_teacher
    ON live_courses (teacher_id)
    WHERE status = 'live';