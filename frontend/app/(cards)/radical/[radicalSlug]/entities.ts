export type Radical = {
  id: string;
  created_at: Date;
  updated_at: Date;
  name: string;
  level: number;
  symbol: Uint8Array;
  meaning_mnemonic: string;
  user_synonyms?: string[];
  user_meaning_note?: string[];
  kanji: RadicalKanji[];
};
export type RadicalKanji = {
  id: string;
  name: string;
  symbol: string;
  reading: string;
  level: number;
};
