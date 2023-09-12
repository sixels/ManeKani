import { Deck } from '../domain/deck';

export interface IDeckRepositoryV1 {
  v1GetDeck(deckId: string): Promise<Deck>;
  v1IsDeckOwner(userID: string, deckId: string): Promise<boolean>;
}
