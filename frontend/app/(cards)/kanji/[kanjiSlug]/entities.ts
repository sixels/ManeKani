export type Kanji = {
  id: string;
  created_at: Date;
  updated_at: Date;
  name: string;
  level: number;
  alt_names: string[];
  symbol: string;
  reading: string;
  onyomi: string[];
  kunyomi: string[];
  nanori: string[];
  meaning_mnemonic: string;
  reading_mnemonic: string;
  user_synonyms: string[];
  user_meaning_note: string;
  user_reading_note: string;
  vocabularies: KanjiVocabulary[];
  radicals: KanjiRadical[];
};
export type KanjiVocabulary = {
  id: string;
  name: string;
  word: string;
  reading: string;
  level: number;
};

export type KanjiRadical = {
  name: string;
  symbol: Uint8Array;
};
