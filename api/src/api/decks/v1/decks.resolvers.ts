import { CreateDeckInput, Deck, UpdateDeckInput } from '@/graphql/deck';
import { DecksAdapter } from 'manekani-core';
import { IDeckRepositoryV1 } from 'manekani-core';
import {
  Args,
  Int,
  Mutation,
  Query,
  ResolveField,
  Resolver,
  Root,
} from '@nestjs/graphql';
import { Authorize, UserData } from '@/api/auth/auth.decorator';
import { Inject } from '@nestjs/common';
import { Subject } from '@/graphql/subject';
import { SubjectsAdapter } from 'manekani-core';
import { ISubjectRepositoryV1 } from 'manekani-core';

@Resolver(Deck)
export class DecksResolver {
  private decksAdapter: DecksAdapter<any>;
  private subjectsAdapter: SubjectsAdapter<any>;

  constructor(
    @Inject('DECKS_REPOSITORY') decksRepository: IDeckRepositoryV1,
    @Inject('SUBJECTS_REPOSITORY') subjectsRepository: ISubjectRepositoryV1,
  ) {
    this.decksAdapter = new DecksAdapter(decksRepository);
    this.subjectsAdapter = new SubjectsAdapter(subjectsRepository);
  }

  @Query(() => Deck)
  async deck(@Args('id') deckId: string): Promise<Deck> {
    return await this.decksAdapter.v1GetDeck(deckId);
  }

  @Query(() => [Deck])
  async decks(
    @Args('page', { type: () => Int, nullable: true }) page?: number,
    @Args('ids', { type: () => [String], nullable: true }) ids?: string[],
    @Args('owners', { type: () => [String], nullable: true }) owners?: string[],
    @Args('content', { type: () => String, nullable: true })
    content?: string,
  ): Promise<Deck[]> {
    return await this.decksAdapter.v1GetDecks({
      page,
      ids,
      owners,
      content,
    });
  }

  @Mutation(() => Deck)
  @Authorize({ scopes: ['deck:create'] })
  async createDeck(
    @UserData('userId') userId: string,
    @Args('data')
    data: CreateDeckInput,
  ): Promise<Deck> {
    return await this.decksAdapter.v1CreateDeck(userId, data);
  }

  @Mutation(() => Deck)
  @Authorize({ scopes: ['deck:update'] })
  async updateDeck(
    @UserData('userId') userId: string,
    @Args('id') deckId: string,
    @Args('data')
    data: UpdateDeckInput,
  ): Promise<Deck> {
    return await this.decksAdapter.v1UpdateDeck(userId, deckId, data);
  }

  @Mutation(() => String)
  @Authorize({ scopes: ['deck:delete'] })
  async deleteDeck(
    @UserData('userId') userId: string,
    @Args('id') deckId: string,
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
