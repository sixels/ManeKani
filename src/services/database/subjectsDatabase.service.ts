import { CastModel, PrismaErrors, inlineAsyncTry, runtimeOmit } from './common';
import {
  CreateSubjectDto,
  GetSubjectsFilters,
  PartialSubject,
  Subject,
  UpdateSubjectDto,
} from '@/core/domain/subject';
import {
  InvalidRequestError,
  ResourceCollidesError,
  ResourceNotFoundError,
} from '@/core/domain/error';

import { ISubjectRepositoryV1 } from '@/core/ports/subjects';
import { Injectable } from '@nestjs/common';
import { PrismaService } from './prisma.service';

@Injectable()
export class SubjectsDatabaseService implements ISubjectRepositoryV1 {
  constructor(private prisma: PrismaService) {}

  private get subjects() {
    return this.prisma.subject;
  }

  async v1GetSubjects(filters: GetSubjectsFilters): Promise<PartialSubject[]> {
    const resultLimit = 500;
    const foundSubjects = await inlineAsyncTry(
      () =>
        this.subjects.findMany({
          where: {
            AND: [
              {
                id: filters.ids && { in: filters.ids },
              },
              {
                category: filters.categories && { in: filters.categories },
              },
              {
                level: filters.levels && { in: filters.levels },
              },
              {
                slug: filters.slugs && { in: filters.slugs },
              },
              {
                deck: {
                  id: filters.decks && { in: filters.decks },
                },
              },
              {
                deck: {
                  ownerId: filters.owners && { in: filters.owners },
                },
              },
            ],
          },
          skip: filters.page && (filters.page - 1) * resultLimit,
          include: {
            ...includeSubjectOwner,
            ...includeSubjectDependencies,
            ...includeSubjectSimilar,
          },
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            description: `An unknown error occurred while retrieving the subjects.`,
            context: { filters },
          },
        });
      },
    );

    return foundSubjects.map((s) => CastModel.intoSubject.partial(s));
  }

  async v1GetSubject(subjectId: string): Promise<Subject | null> {
    const foundSubject = await inlineAsyncTry(
      () =>
        this.subjects.findUnique({
          where: { id: subjectId },
          include: {
            ...includeSubjectOwner,
            ...includeSubjectDependencies,
            ...includeSubjectSimilar,
          },
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { subjectId },
            description: `An unknown error occurred while retrieving the subject with id "${subjectId}".`,
          },
          byError: {
            [PrismaErrors.NotFound]: [
              ResourceNotFoundError,
              {
                context: { subjectId },
                description: `No subjects with id "${subjectId}" were found.`,
              },
            ],
          },
        });
      },
    );

    return foundSubject && CastModel.intoSubject(foundSubject);
  }

  async v1CreateSubject(
    userId: string,
    subject: CreateSubjectDto,
  ): Promise<Subject> {
    const createdSubject = await inlineAsyncTry(
      () =>
        this.subjects.create({
          data: {
            ...runtimeOmit(subject, ['dependencies', 'similar', 'dependents']),
            similarTo: {
              connect: (subject.similar ?? []).map((id) => ({ id })),
            },
            dependsOn: {
              connect: (subject.dependencies ?? []).map((id) => ({ id })),
            },
            dependedBy: {
              connect: (subject.dependents ?? []).map((id) => ({ id })),
            },
          },
          include: includeSubjectOwner,
        }),
      (error) => {
        const errorContext = {
          name: subject.name,
          category: subject.category,
          slug: subject.slug,
          value: subject.value,
          valueImage: subject.valueImage,
          deckId: subject.deckId,
          similar: subject.similar,
          dependencies: subject.dependencies,
          dependents: subject.dependents,
          ownerId: userId,
        };

        throw PrismaErrors.match(error, {
          fallback: {
            description: `An unknown error occurred while creating the subject.`,
            context: errorContext,
          },
          byError: {
            [PrismaErrors.UniqueConstraint]: [
              ResourceCollidesError,
              {
                context: errorContext,
                description: `A subject with same slug and category already already exists in this deck.`,
              },
            ],
            [PrismaErrors.ForeignKeyConstraint]: [
              InvalidRequestError,
              {
                context: errorContext,
                description: `Could not find a deck with id "${subject.deckId}".`,
              },
            ],
            [PrismaErrors.DependencyNotFound]: [
              InvalidRequestError,
              {
                context: errorContext,
                description: `The subject depends on a subject which does not exist. check if all dependencies, dependents or similar subjects exist.`,
              },
            ],
          },
        });
      },
    );

    return CastModel.intoSubject(createdSubject);
  }

  async v1UpdateSubject(
    userId: string,
    subjectId: string,
    changes: UpdateSubjectDto,
  ): Promise<Subject> {
    const updatedSubject = await inlineAsyncTry(
      () =>
        this.prisma.subject.update({
          where: { id: subjectId, deck: { ownerId: userId } },
          data: {
            ...runtimeOmit(changes, ['similar', 'dependencies', 'dependents']),
            similarTo:
              changes.similar !== undefined
                ? { connect: changes.similar.map((id) => ({ id })) }
                : undefined,
            dependsOn:
              changes.dependencies !== undefined
                ? { connect: changes.dependencies.map((id) => ({ id })) }
                : undefined,
            dependedBy:
              changes.dependents !== undefined
                ? { connect: changes.dependents.map((id) => ({ id })) }
                : undefined,
          },
          include: includeSubjectOwner,
        }),
      (error) => {
        const errorContext = {
          name: changes.name,
          category: changes.category,
          slug: changes.slug,
          value: changes.value,
          valueImage: changes.valueImage,
          similar: changes.similar,
          dependencies: changes.dependencies,
          dependents: changes.dependents,
        };
        throw PrismaErrors.match(error, {
          fallback: {
            description:
              'An unknown error occurred while updating the subject.',
            context: errorContext,
          },
          byError: {
            [PrismaErrors.UniqueConstraint]: [
              ResourceCollidesError,
              {
                context: errorContext,
                description: `A subject with same slug and category already already exists in this deck.`,
              },
            ],
            [PrismaErrors.DependencyNotFound]: [
              InvalidRequestError,
              {
                context: errorContext,
                description: `The subject or one of its dependencies does not exist. check if all dependencies, dependents or similar subjects exist.`,
              },
            ],
          },
        });
      },
    );

    return CastModel.intoSubject(updatedSubject);
  }

  async v1DeleteSubject(userId: string, subjectId: string): Promise<void> {
    const _deletedSubject = await inlineAsyncTry(
      () =>
        this.subjects.delete({
          where: { id: subjectId, deck: { ownerId: userId } },
          select: {},
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            description: `An unknown error occurred while deleting the subject with id "${subjectId}".`,
            context: { subjectId },
          },
          byError: {
            [PrismaErrors.NotFound]: [
              ResourceNotFoundError,
              {
                context: { subjectId },
                description: `No subjects with id "${subjectId}" were found.`,
              },
            ],
          },
        });
      },
    );
  }
}

const includeSubjectOwner = { deck: { select: { ownerId: true } } };
const includeSubjectDependencies = {
  dependsOn: { select: { id: true } },
  dependedBy: { select: { id: true } },
};
const includeSubjectSimilar = {
  similarFrom: { select: { id: true } },
  similarTo: { select: { id: true } },
};
