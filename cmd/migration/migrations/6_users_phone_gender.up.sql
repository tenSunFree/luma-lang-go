ALTER TABLE users
    ADD COLUMN IF NOT EXISTS full_name VARCHAR(100),
    ADD COLUMN IF NOT EXISTS phone     VARCHAR(20) UNIQUE,
    ADD COLUMN IF NOT EXISTS gender    VARCHAR(10)
        CHECK (gender IN ('male', 'female', 'other'));

CREATE INDEX IF NOT EXISTS idx_users_phone ON users (phone);