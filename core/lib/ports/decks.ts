import {
	CreateDeckDto,
	Deck,
	GetDecksFilters,
	UpdateDeckDto,
} from "../domain/deck";

export interface IDeckRepositoryV1 {
	v1GetDecks(filters: GetDecksFilters): Promise<Deck[]>;
	v1GetDeck(deckId: string): Promise<Deck | null>;
	v1CreateDeck(userId: string, deck: CreateDeckDto): Promise<Deck>;
	v1UpdateDeck(
		userId: string,
		deckId: string,
		changes: UpdateDeckDto,
	): Promise<Deck>;
	v1DeleteDeck(userId: string, deckId: string): Promise<void>;
	v1IsDeckOwner(userID: string, deckId: string): Promise<boolean>;
}
