CREATE TABLE radicals (
    id uuid UNIQUE DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    name TEXT UNIQUE NOT NULL,
    user_synonyms TEXT [],
    symbol VARCHAR(5) NOT NULL UNIQUE,
    meaning_mnemonic TEXT NOT NULL,
    user_meaning_note TEXT,

    PRIMARY KEY (id, name)
);

-- update `updated_at` automatically

CREATE TRIGGER update_timestamp_trigger
BEFORE UPDATE ON radicals
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();