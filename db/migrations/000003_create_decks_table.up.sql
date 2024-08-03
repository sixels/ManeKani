CREATE TABLE IF NOT EXISTS decks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    slug TEXT NOT NULL,
    image TEXT,
    tags TEXT [],
    is_featured BOOLEAN DEFAULT FALSE NOT NULL,
    is_public BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    created_by_user_id UUID NOT NULL
);