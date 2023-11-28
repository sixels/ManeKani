import { CreateDeckDto, Deck, GetDecksFilters } from "@manekani/core";
import { ResourceCollidesError, ResourceNotFoundError } from "@manekani/core";
import { CastModel, PrismaErrors, inlineAsyncTry } from "./common";

import { IDeckRepositoryV1 } from "@manekani/core";
import { PrismaClient } from "@prisma/client";

export const DecksProviderLabel = "DECKS_REPOSITORY";

export class DecksDatabase implements IDeckRepositoryV1 {
	constructor(private client: PrismaClient) {}

	private get decks() {
		return this.client.deck;
	}

	async v1GetDecks(filters: GetDecksFilters): Promise<Deck[]> {
		const resultLimit = 100;
		const foundDecks = await inlineAsyncTry(
			() =>
				this.decks.findMany({
					where: {
						AND: [
							{ id: filters.ids && { in: filters.ids } },
							{ ownerId: filters.owners && { in: filters.owners } },
							{ name: filters.content && { contains: filters.content } },
							{ description: filters.content && { contains: filters.content } },
						],
					},
					skip: filters.page && (filters.page - 1) * resultLimit,
					include: {
						ownedBy: { select: { id: true } },
						subjects: { select: { id: true } },
					},
				}),
			(error) => {
				throw PrismaErrors.match(error, {
					fallback: {
						description:
							"An unknown error occurred while retrieving the decks.",
						context: { filters },
					},
				});
			},
		);

		return foundDecks.map((d) => CastModel.intoDeck(d));
	}
	async v1GetDeck(deckId: string): Promise<Deck | null> {
		const foundDeck = await inlineAsyncTry(
			() =>
				this.decks.findUnique({
					where: { id: deckId },
					include: {
						ownedBy: { select: { id: true } },
						subjects: { select: { id: true } },
					},
				}),
			(error) => {
				throw PrismaErrors.match(error, {
					fallback: {
						context: { deckId },
						description: `An unknown error occurred while retrieving the deck with id "${deckId}".`,
					},
					byError: {
						[PrismaErrors.NotFound]: [
							ResourceNotFoundError,
							{
								context: { deckId },
								description: `No deck with id "${deckId}" were found.`,
							},
						],
					},
				});
			},
		);
		return foundDeck && CastModel.intoDeck(foundDeck);
	}

	async v1CreateDeck(userId: string, deck: CreateDeckDto): Promise<Deck> {
		const createdDeck = await inlineAsyncTry(
			() =>
				this.decks.create({
					data: {
						...deck,
						ownedBy: { connect: { id: userId } },
					},
					include: {
						ownedBy: { select: { id: true } },
						subjects: { select: { id: true } },
					},
				}),
			(error) => {
				throw PrismaErrors.match(error, {
					fallback: {
						description: "An unknown error occurred while creating the deck.",
						context: { deck, userId },
					},
					byError: {
						[PrismaErrors.UniqueConstraint]: [
							ResourceCollidesError,
							{
								description: "You already have a deck with the same name.",
								context: { deck, userId },
							},
						],
						[PrismaErrors.ForeignKeyConstraint]: [
							ResourceNotFoundError,
							{
								description: "The deck owner does not exist.",
								context: { deck, userId },
							},
						],
					},
				});
			},
		);

		return CastModel.intoDeck(createdDeck);
	}
	async v1UpdateDeck(
		userId: string,
		deckId: string,
		changes: {
			name?: string | undefined;
			description?: string | undefined;
			subjectIds?: string[] | undefined;
		},
	): Promise<Deck> {
		const updatedDeck = await inlineAsyncTry(
			() =>
				this.decks.update({
					where: { id: deckId, ownerId: userId },
					data: {
						...changes,
					},
					include: {
						ownedBy: { select: { id: true } },
						subjects: { select: { id: true } },
					},
				}),
			(error) => {
				throw PrismaErrors.match(error, {
					fallback: {
						description: "An unknown error occurred while updating the deck.",
						context: { deckId, changes, userId },
					},
					byError: {
						[PrismaErrors.UniqueConstraint]: [
							ResourceCollidesError,
							{
								description: "You already have a deck with the same name.",
								context: { deckId, changes, userId },
							},
						],
						[PrismaErrors.ForeignKeyConstraint]: [
							ResourceNotFoundError,
							{
								description: "The deck owner does not exist.",
								context: { deckId, changes, userId },
							},
						],
					},
				});
			},
		);

		return CastModel.intoDeck(updatedDeck);
	}
	v1DeleteDeck(_userId: string, _deckId: string): Promise<void> {
		throw new Error("Method not implemented.");
	}

	v1IsDeckOwner(_userID: string, _deckId: string): Promise<boolean> {
		throw new Error("Method not implemented.");
	}
}
