import {
	Args,
	Int,
	Mutation,
	Query,
	ResolveField,
	Resolver,
	Root,
} from "@nestjs/graphql";

import { Authorize, UserData } from "@/api/auth/auth.decorator";
import { Deck } from "@/graphql/deck";
import {
	CreateSubjectInput,
	Subject,
	UpdateSubjectInput,
} from "@/graphql/subject";
import { DecksProviderLabel } from "@/infra/database/decksDatabase.service";
import { SubjectsProviderLabel } from "@/infra/database/subjectsDatabase.service";
import {
	DecksAdapter,
	FilesAdapter,
	IDeckRepositoryV1,
	IFilesRepositoryV1,
	ISubjectRepositoryV1,
	SubjectsAdapter,
} from "@manekani/core";
import { Inject } from "@nestjs/common";
import { FileStorageProviderLabel } from "@/infra/files/files.service";

@Resolver(Subject)
export class SubjectsResolver {
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
	private subjectsAdapter: SubjectsAdapter<any>;
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
	private decksAdapter: DecksAdapter<any>;
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
	private filesAdapter: FilesAdapter<any>;

	constructor(
		@Inject(SubjectsProviderLabel) subjectsRepository: ISubjectRepositoryV1,
		@Inject(DecksProviderLabel) decksRepository: IDeckRepositoryV1,
		@Inject(FileStorageProviderLabel) filesRepository: IFilesRepositoryV1,
	) {
		this.decksAdapter = new DecksAdapter(decksRepository);
		this.filesAdapter = new FilesAdapter(filesRepository);
		this.subjectsAdapter = new SubjectsAdapter(
			subjectsRepository,
		).withFilesAdapter(filesRepository);
	}

	@Query(() => Subject)
	async subject(@Args("id") subjectId: string): Promise<Subject> {
		return await this.subjectsAdapter.v1GetSubject(subjectId);
	}

	@Query(() => [Subject])
	async subjects(
		@Args("decks", { type: () => [String], nullable: false })
		decks: string[],
		@Args("page", { type: () => Int, nullable: true })
		page?: number,
		@Args("ids", { type: () => [String], nullable: true })
		ids?: string[],
		@Args("categories", { type: () => [String], nullable: true })
		categories?: string[],
		@Args("levels", { type: () => [Int], nullable: true })
		levels?: number[],
		@Args("slugs", { type: () => [String], nullable: true })
		search?: string[],
		@Args("owners", { type: () => [String], nullable: true })
		owners?: string[],
	): Promise<Subject[]> {
		if (!decks.length) {
			return [];
		}

		return await this.subjectsAdapter.v1GetSubjects({
			page,
			ids,
			categories,
			levels,
			slugs: search,
			decks,
			owners,
		});
	}

	@Mutation(() => Subject)
	@Authorize({ scopes: ["subject:create"] })
	async createSubject(
		@UserData("userId") userId: string,
		@Args("deckId") deckId: string,
		@Args("data") data: CreateSubjectInput,
	): Promise<Subject> {
		return await this.subjectsAdapter.v1CreateSubject(userId, deckId, data);
	}

	@Mutation(() => Subject)
	@Authorize({ scopes: ["subject:update"] })
	async updateSubject(
		@UserData("userId") userId: string,
		@Args("subjectId") subjectId: string,
		@Args("data") data: UpdateSubjectInput,
	): Promise<Subject> {
		return await this.subjectsAdapter.v1UpdateSubject(userId, subjectId, data);
	}

	@Mutation(() => String)
	@Authorize({ scopes: ["subject:delete"] })
	async deleteSubject(
		@UserData("userId") userId: string,
		@Args("subjectId") subjectId: string,
	): Promise<string> {
		await this.subjectsAdapter.v1DeleteSubject(userId, subjectId);
		return `Deleted subject ${subjectId}`;
	}

	@ResolveField(() => Deck)
	async deck(@Root() subject: Subject): Promise<Deck> {
		return await this.decksAdapter.v1GetDeck(subject.deckId);
	}
}
