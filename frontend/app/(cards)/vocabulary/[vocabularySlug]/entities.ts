export type Vocabulary = {
  id: string;
  created_at: Date;
  updated_at: Date;
  name: string;
  level: number;
  alt_names: string[];
  word: string;
  word_type: string[];
  reading: string;
  meaning_mnemonic: string;
  reading_mnemonic: string;
  user_synonyms?: string[];
  user_meaning_note?: string;
  user_reading_note?: string;
};
