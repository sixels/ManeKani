export interface CreateKanjiRequest {
  alt_names?: string[];
  kunyomi?: string[];
  level?: number;
  meaning_mnemonic?: string;
  name?: string;
  nanori?: string[];
  onyomi?: string[];
  radical_composition?: string[];
  reading?: string;
  reading_mnemonic?: string;
  similar?: string[];
  symbol?: string;
}

export interface CreateRadicalRequest {
  level?: number;
  meaning_mnemonic?: string;
  name?: string;
  symbol?: string;
  symbol_image?: number[];
}

export interface CreateVocabularyRequest {
  alt_names?: string[];
  alt_readings?: string[];
  kanji_composition?: string[];
  level?: number;
  meaning_mnemonic?: string;
  name?: string;
  patterns?: Pattern[];
  reading?: string;
  reading_mnemonic?: string;
  sentences?: Sentence[];
  word?: string;
  word_type?: string[];
}

export interface Kanji {
  alt_names?: string[];
  created_at?: string;
  id?: string;
  kunyomi: string[];
  level: number;
  meaning_mnemonic: string;
  name: string;
  nanori: string[];
  onyomi: string[];
  reading: string;
  reading_mnemonic: string;
  similar?: string[];
  symbol: string;
  updated_at?: string;
}

export interface Level {
  kanji?: PartialKanjiResponse[];
  radical?: PartialRadicalResponse[];
  vocabulary?: PartialVocabularyResponse[];
}

export interface PartialKanjiResponse {
  alt_names?: string[];
  id?: string;
  level?: number;
  name?: string;
  reading?: string;
  symbol?: string;
}

export interface PartialRadicalResponse {
  id: string;
  level: number;
  name: string;
  symbol: string;
}

export interface PartialVocabularyResponse {
  alt_names?: string[];
  id: string;
  level: number;
  name: string;
  reading: string;
  word: string;
}

export interface Radical {
  created_at?: string;
  id: string;
  level: number;
  meaning_mnemonic: string;
  name: string;
  symbol: string;
  updated_at?: string;
}

export interface UpdateKanjiRequest {
  alt_names?: string[];
  kunyomi?: string[];
  level?: number;
  meaning_mnemonic?: string;
  name?: string;
  nanori?: string[];
  onyomi?: string[];
  radical_composition?: string[];
  reading?: string;
  reading_mnemonic?: string;
  similar?: string[];
}

export interface UpdateRadicalRequest {
  level?: number;
  meaning_mnemonic?: string;
  symbol?: string;
}

export interface UpdateVocabularyRequest {
  alt_names?: string[];
  alt_readings?: string[];
  kanji_composition?: string[];
  level?: number;
  meaning_mnemonic?: string;
  name?: string;
  patterns?: Pattern[];
  reading?: string;
  reading_mnemonic?: string;
  sentences?: Sentence[];
  word_type?: string[];
}

export interface Vocabulary {
  alt_names?: string[];
  alt_readings?: string[];
  created_at?: string;
  id: string;
  level: number;
  meaning_mnemonic: string;
  name: string;
  patterns?: Pattern[];
  reading: string;
  reading_mnemonic: string;
  sentences?: Sentence[];
  updated_at?: string;
  word: string;
  word_type: string[];
}

export interface Pattern {
  name: string;
  sentences: Sentence[];
}

export interface Sentence {
  meaning: string;
  sentence: string;
}
