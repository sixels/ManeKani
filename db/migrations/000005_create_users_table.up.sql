CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL,
    username TEXT,
    display_name TEXT,
    is_verified BOOLEAN DEFAULT FALSE NOT NULL,
    is_complete BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    CONSTRAINT user_email_unique UNIQUE (email),
    CONSTRAINT user_username_unique UNIQUE (username)
);