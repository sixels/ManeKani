import {
  Args,
  Int,
  Mutation,
  Query,
  ResolveField,
  Resolver,
  Root,
} from '@nestjs/graphql';

import {
  CreateSubjectInput,
  Subject,
  UpdateSubjectInput,
} from '@/graphql/subject';
import { Authorize, UserData } from '@/api/auth/auth.decorator';
import { SubjectsAdapter } from 'manekani-core';
import { ISubjectRepositoryV1 } from 'manekani-core';
import { Inject } from '@nestjs/common';
import { Deck } from '@/graphql/deck';
import { DecksAdapter } from 'manekani-core';
import { IDeckRepositoryV1 } from 'manekani-core';

@Resolver(Subject)
export class SubjectsResolver {
  private subjectsAdapter: SubjectsAdapter<any>;
  private decksAdapter: DecksAdapter<any>;

  constructor(
    @Inject('SUBJECTS_REPOSITORY') subjectsRepository: ISubjectRepositoryV1,
    @Inject('DECKS_REPOSITORY') decksRepository: IDeckRepositoryV1,
  ) {
    this.subjectsAdapter = new SubjectsAdapter(subjectsRepository);
    this.decksAdapter = new DecksAdapter(decksRepository);
  }

  @Query(() => Subject)
  async subject(@Args('id') subjectId: string): Promise<Subject> {
    return await this.subjectsAdapter.v1GetSubject(subjectId);
  }

  @Query(() => [Subject])
  async subjects(
    @Args('page', { type: () => Int, nullable: true })
    page?: number,
    @Args('ids', { type: () => [String], nullable: true })
    ids?: string[],
    @Args('categories', { type: () => [String], nullable: true })
    categories?: string[],
    @Args('levels', { type: () => [Int], nullable: true })
    levels?: number[],
    @Args('slugs', { type: () => [String], nullable: true })
    search?: string[],
    @Args('decks', { type: () => [String], nullable: true })
    decks?: string[],
    @Args('owners', { type: () => [String], nullable: true })
    owners?: string[],
  ): Promise<Subject[]> {
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
  @Authorize({ scopes: ['subject:create'] })
  async createSubject(
    @UserData('userId') userId: string,
    @Args('deckId') deckId: string,
    @Args('data') data: CreateSubjectInput,
  ): Promise<Subject> {
    return await this.subjectsAdapter.v1CreateSubject(userId, deckId, data);
  }

  @Mutation(() => Subject)
  @Authorize({ scopes: ['subject:update'] })
  async updateSubject(
    @UserData('userId') userId: string,
    @Args('subjectId') subjectId: string,
    @Args('data') data: UpdateSubjectInput,
  ): Promise<Subject> {
    return await this.subjectsAdapter.v1UpdateSubject(userId, subjectId, data);
  }

  @Mutation(() => String)
  @Authorize({ scopes: ['subject:delete'] })
  async deleteSubject(
    @UserData('userId') userId: string,
    @Args('subjectId') subjectId: string,
  ): Promise<string> {
    await this.subjectsAdapter.v1DeleteSubject(userId, subjectId);
    return `Deleted subject ${subjectId}`;
  }

  @ResolveField(() => Deck)
  async deck(@Root() subject: Subject): Promise<Deck> {
    return await this.decksAdapter.v1GetDeck(subject.deckId);
  }
}
