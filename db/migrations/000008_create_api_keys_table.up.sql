CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    token TEXT NOT NULL,
    prefix TEXT NOT NULL,
    claims JSONB NOT NULL,
    revoked_at TIMESTAMPTZ,
    used_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    created_by_user_id UUID NOT NULL,
    CONSTRAINT unique_name_created_by_user UNIQUE (name, created_by_user_id),
    CONSTRAINT created_by_user_id_fk FOREIGN KEY (created_by_user_id) REFERENCES users(id) ON DELETE CASCADE
);