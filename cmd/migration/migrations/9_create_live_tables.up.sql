CREATE TABLE IF NOT EXISTS live_courses (
    id                  VARCHAR(100) PRIMARY KEY,
    course_id           VARCHAR(100) NOT NULL,
    title               VARCHAR(255) NOT NULL,
    category            VARCHAR(100) NOT NULL DEFAULT '',
    level               VARCHAR(50)  NOT NULL DEFAULT '',
    course_type         VARCHAR(30)  NOT NULL DEFAULT 'required',
    status              VARCHAR(30)  NOT NULL DEFAULT 'scheduled'
        CHECK (status IN ('scheduled', 'live', 'ended', 'cancelled')),

    scheduled_start_at  TIMESTAMPTZ NOT NULL,
    started_at          TIMESTAMPTZ,
    ended_at            TIMESTAMPTZ,

    teacher_id          VARCHAR(100) NOT NULL,
    teacher_name        VARCHAR(100) NOT NULL DEFAULT '',
    teacher_avatar_url  TEXT,

    thumbnail_url       TEXT,
    textbook_url        TEXT,

    agora_channel_name  VARCHAR(100) NOT NULL UNIQUE,
    teacher_camera_uid  INTEGER NOT NULL DEFAULT 1000,
    teacher_screen_uid  INTEGER NOT NULL DEFAULT 2000,

    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_live_courses_status ON live_courses(status);
CREATE INDEX IF NOT EXISTS idx_live_courses_scheduled_start_at ON live_courses(scheduled_start_at);

CREATE TABLE IF NOT EXISTS live_participants (
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    live_id      VARCHAR(100) NOT NULL REFERENCES live_courses(id) ON DELETE CASCADE,
    user_id      VARCHAR(100) NOT NULL,
    display_name VARCHAR(100) NOT NULL DEFAULT '',
    avatar_url   TEXT,
    agora_uid    INTEGER      NOT NULL,
    role         VARCHAR(30)  NOT NULL DEFAULT 'audience'
        CHECK (role IN ('teacher', 'audience')),
    joined_at    TIMESTAMPTZ  NOT NULL DEFAULT now(),
    left_at      TIMESTAMPTZ,
    last_seen_at TIMESTAMPTZ  NOT NULL DEFAULT now(),

    UNIQUE(live_id, user_id, agora_uid)
);

CREATE INDEX IF NOT EXISTS idx_live_participants_live_id ON live_participants(live_id);
CREATE INDEX IF NOT EXISTS idx_live_participants_online ON live_participants(live_id, left_at);

CREATE TABLE IF NOT EXISTS live_reminders (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    live_id    VARCHAR(100) NOT NULL REFERENCES live_courses(id) ON DELETE CASCADE,
    user_id    VARCHAR(100) NOT NULL,
    remind_at  TIMESTAMPTZ  NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT now(),

    UNIQUE(live_id, user_id)
);

CREATE TABLE IF NOT EXISTS live_chat_messages (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    live_id    VARCHAR(100) NOT NULL REFERENCES live_courses(id) ON DELETE CASCADE,
    user_id    VARCHAR(100) NOT NULL,
    message    TEXT         NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_live_chat_messages_live_id_created_at
    ON live_chat_messages(live_id, created_at);