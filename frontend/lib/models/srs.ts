import { MinimalSubject, PartialSubject } from "./subject";

export interface SRSUserData {
  cards: Card[];
}

export interface Card {
  id: string;
  created_at: string;
  updated_at: string;
  subject: PartialSubject;
  progress: number;
  total_errors: number;
  unlocked_at?: string;
  started_at?: string;
  passed_at?: string;
  available_at?: string;
  burned_at?: string;
}

export interface SessionQueue {
  cards: QueueItem[];
}

export interface QueueItem {
  card_id: string;
  answers: QueueItemAnswers[];
  subject: MinimalSubject;
}

export interface QueueItemAnswers {
  study_item_type: string;
  expected: string[];
  blacklisted: string[];
}
