CREATE TABLE kanjis (
    id uuid UNIQUE DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    name TEXT UNIQUE,
    alt_names TEXT [] NOT NULL,
    user_synonyms TEXT [],
    symbol VARCHAR(5) NOT NULL UNIQUE,
    reading TEXT NOT NULL,
    onyomi TEXT [] NOT NULL,
    kunyomi TEXT [] NOT NULL,
    nanori TEXT [] NOT NULL,
    meaning_mnemonic TEXT NOT NULL,
    reading_mnemonic TEXT NOT NULL,
    user_meaning_note TEXT,
    user_reading_note TEXT,

    PRIMARY KEY (id, name)
);

CREATE TABLE kanjis_radicals (
    kanji_id uuid,
    radical_name TEXT,
    PRIMARY KEY (kanji_id, radical_name),
    CONSTRAINT fk_kanji FOREIGN KEY(kanji_id) REFERENCES kanjis(id),
    CONSTRAINT fk_radical FOREIGN KEY(radical_name) REFERENCES radicals(name)
);

-- update `updated_at` automatically

CREATE TRIGGER update_timestamp_trigger
BEFORE UPDATE ON kanjis
FOR EACH ROW
EXECUTE PROCEDURE update_timestamp();