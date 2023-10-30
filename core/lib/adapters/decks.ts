import {
  CreateDeckDto,
  CreateDeckSchema,
  Deck,
  GetDecksFilters,
  GetDecksFiltersSchema,
  UpdateDeckDto,
  UpdateDeckSchema,
} from '../domain/deck';

import { IDeckRepositoryV1 } from '../ports/decks';
import { ResourceNotFoundError } from '../domain/error';
import { Validator } from '../validator';
import { validateId } from './common';

export const CreateDeckValidator = new Validator(CreateDeckSchema);
export const UpdateDeckValidator = new Validator(UpdateDeckSchema);
export const DecksFiltersValidator = new Validator(GetDecksFiltersSchema);

export class DecksAdapter<R extends IDeckRepositoryV1> {
  constructor(private decksRepository: R) {}

  v1GetDecks(filters: GetDecksFilters): Promise<Deck[]> {
    DecksFiltersValidator.validate(filters);
    return this.decksRepository.v1GetDecks(filters);
  }

  async v1GetDeck(deckId: string): Promise<Deck> {
    validateId(deckId);

    const foundDeck = await this.decksRepository.v1GetDeck(deckId);
    if (!foundDeck) {
      throw new ResourceNotFoundError({
        cause: new Error('Deck not found'),
        context: { deckId },
        description: `No decks with id "${deckId}" were found.`,
      });
    }
    return foundDeck;
  }

  v1CreateDeck(userId: string, deck: CreateDeckDto): Promise<Deck> {
    CreateDeckValidator.validate(deck);

    console.debug('creating deck:', { userId, deck });
    return this.decksRepository.v1CreateDeck(userId, deck);
  }

  v1UpdateDeck(
    userId: string,
    deckId: string,
    deck: UpdateDeckDto,
  ): Promise<Deck> {
    validateId(deckId);
    UpdateDeckValidator.validate(deck);

    console.debug('updating deck:', { deckId, deck });
    return this.decksRepository.v1UpdateDeck(userId, deckId, deck);
  }

  v1DeleteDeck(userId: string, deckId: string): Promise<void> {
    validateId(deckId);

    console.debug('deleting deck:', { deckId });
    return this.decksRepository.v1DeleteDeck(userId, deckId);
  }
}
