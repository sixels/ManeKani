CREATE TABLE IF NOT EXISTS subjects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category TEXT NOT NULL,
    level INT NOT NULL,
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    priority INT NOT NULL,
    data JSONB NOT NULL,
    study_data JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    deck_id UUID NOT NULL,
    CONSTRAINT subject_category_slug_deck_id UNIQUE (category, slug, deck_id)
);