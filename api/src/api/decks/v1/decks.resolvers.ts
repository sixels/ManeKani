import { Authorize, UserData } from "@/api/auth/auth.decorator";
import { CreateDeckInput, Deck, UpdateDeckInput } from "@/graphql/deck";
import { Subject } from "@/graphql/subject";
import {
	DecksAdapter,
	IDeckRepositoryV1,
	ISubjectRepositoryV1,
	SubjectsAdapter,
} from "@manekani/core";
import { Inject } from "@nestjs/common";
import {
	Args,
	Int,
	Mutation,
	Query,
	ResolveField,
	Resolver,
	Root,
} from "@nestjs/graphql";

@Resolver(Deck)
export class DecksResolver {
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
	private decksAdapter: DecksAdapter<any>;
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
	private subjectsAdapter: SubjectsAdapter<any>;

	constructor(
		@Inject("DECKS_REPOSITORY") decksRepository: IDeckRepositoryV1,
		@Inject("SUBJECTS_REPOSITORY") subjectsRepository: ISubjectRepositoryV1,
	) {
		this.decksAdapter = new DecksAdapter(decksRepository);
		this.subjectsAdapter = new SubjectsAdapter(subjectsRepository);
	}

	@Query(() => Deck)
	async deck(@Args("id") deckId: string): Promise<Deck> {
		return await this.decksAdapter.v1GetDeck(deckId);
	}

	@Query(() => [Deck])
	async decks(
		@Args("page", { type: () => Int, nullable: true }) page?: number,
		@Args("ids", { type: () => [String], nullable: true }) ids?: string[],
		@Args("owners", { type: () => [String], nullable: true }) owners?: string[],
		@Args("names", { type: () => [String], nullable: true }) names?: string[],
		@Args("content", { type: () => String, nullable: true })
		content?: string,
	): Promise<Deck[]> {
		return await this.decksAdapter.v1GetDecks({
			page,
			ids,
			owners,
			names,
			content,
		});
	}

	@Query(() => Deck)
	@Authorize()
	async deckByName(@Args("name") deckName: string): Promise<Deck> {
		// TODO: we should have a way to get a single deck by name implemented at the core
		const decks = await this.decksAdapter.v1GetDecks({ names: [deckName] });
		if (decks.length === 0) {
			throw new Error(`Deck ${deckName} not found`);
		}
		return decks[0];
	}

	@Mutation(() => Deck)
	@Authorize({ scopes: ["deck:create"] })
	async createDeck(
		@UserData("userId") userId: string,
		@Args("data")
		data: CreateDeckInput,
	): Promise<Deck> {
		return await this.decksAdapter.v1CreateDeck(userId, data);
	}

	@Mutation(() => Deck)
	@Authorize({ scopes: ["deck:update"] })
	async updateDeck(
		@UserData("userId") userId: string,
		@Args("id") deckId: string,
		@Args("data")
		data: UpdateDeckInput,
	): Promise<Deck> {
		return await this.decksAdapter.v1UpdateDeck(userId, deckId, data);
	}

	@Mutation(() => String)
	@Authorize({ scopes: ["deck:delete"] })
	async deleteDeck(
		@UserData("userId") userId: string,
		@Args("id") deckId: string,
	): Promise<string> {
		await this.decksAdapter.v1DeleteDeck(userId, deckId);
		return `Deck ${deckId} deleted`;
	}

	@ResolveField(() => [Subject])
	async subjects(@Root() deck: Deck): Promise<Subject[]> {
		return await this.subjectsAdapter.v1GetSubjects({
			decks: [deck.id],
			ids: deck.subjectIds,
		});
	}
}
