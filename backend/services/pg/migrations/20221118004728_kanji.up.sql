CREATE TABLE kanji (
    id uuid UNIQUE DEFAULT uuid_generate_v4 (),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL,
    level INT NOT NULL,
    alt_names TEXT [] NOT NULL,
    user_synonyms TEXT [],
    symbol VARCHAR(1) UNIQUE NOT NULL,
    reading TEXT NOT NULL,
    onyomi TEXT [] NOT NULL,
    kunyomi TEXT [] NOT NULL,
    nanori TEXT [] NOT NULL,
    meaning_mnemonic TEXT NOT NULL,
    reading_mnemonic TEXT NOT NULL,
    user_meaning_note TEXT,
    user_reading_note TEXT,
    PRIMARY KEY (id, symbol)
);
CREATE TABLE kanji_radicals (
    kanji_symbol VARCHAR(1),
    radical_name TEXT,
    PRIMARY KEY (kanji_symbol, radical_name),
    CONSTRAINT fk_kanji FOREIGN KEY(kanji_symbol) REFERENCES kanji(symbol),
    CONSTRAINT fk_radical FOREIGN KEY(radical_name) REFERENCES radicals(name)
);
-- update `updated_at` automatically
CREATE TRIGGER update_timestamp_trigger BEFORE
UPDATE ON kanji FOR EACH ROW EXECUTE PROCEDURE update_timestamp();