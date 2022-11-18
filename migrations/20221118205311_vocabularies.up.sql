CREATE TABLE vocabularies (
    id uuid UNIQUE DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    name TEXT,
    alt_names TEXT [] NOT NULL,
    user_synonyms TEXT [],
    word TEXT NOT NULL UNIQUE,
    word_type TEXT [] NOT NULL,
    reading TEXT NOT NULL,
    meaning_mnemonic TEXT NOT NULL,
    reading_mnemonic TEXT NOT NULL,
    user_meaning_note TEXT,
    user_reading_note TEXT,

    PRIMARY KEY (id, name)
);

-- TODO: CASCADE on UPDATE/DELETE
CREATE TABLE vocabularies_kanjis (
    vocabulary_id uuid,
    kanji_id uuid,
    PRIMARY KEY (vocabulary_id, kanji_id),
    CONSTRAINT fk_vocabulary FOREIGN KEY(vocabulary_id) REFERENCES vocabularies(id)
    CONSTRAINT fk_kanji FOREIGN KEY(kanji_id) REFERENCES kanjis(id),
);

-- update `updated_at` automatically

CREATE TRIGGER update_timestamp_trigger
BEFORE UPDATE ON vocabularies
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();