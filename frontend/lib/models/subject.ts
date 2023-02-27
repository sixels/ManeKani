export interface Subject {
  id: string;
  created_at: string;
  updated_at: string;
  kind: string;
  level: number;
  name: string;
  value?: string;
  value_image?: RemoteContent;
  slug: string;
  priority: number;
  resources?: { [key: string]: RemoteContent[] };
  complimentary_study_data?: { [key: string]: any };
  study_data: StudyData[];
  dependencies: string[];
  dependents: string[];
  similars: string[];
  deck: string;
  owner: string;
}

export interface PartialSubject {
  id: string;
  kind: string;
  level: number;
  name: string;
  value?: string;
  value_image?: RemoteContent;
  slug: string;
  priority: number;
  study_data: StudyData[];
  dependencies: string[];
  dependents: string[];
  similars: string[];
  deck: string;
  owner: string;
}

export type MinimalSubject = Omit<
  PartialSubject,
  "study_data" | "dependencies" | "dependents" | "similars" | "owners"
>;

export interface RemoteContent {
  url: string;
  content_type: string;
  metadata?: { [key: string]: any };
}

export interface StudyData {
  kind: string;
  items: StudyItem[];
  mnemonic?: string;
}

export interface StudyItem {
  value: string;
  is_primary: boolean;
  is_valid_answer: boolean;
  is_hidden: boolean;
  category?: string;
}

export interface Patterns {
  [key: string]: Sentence[];
}

export interface Sentence {
  en: string;
  jp: string;
}
