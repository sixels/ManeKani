import { Deck, Token } from 'manekani-core';
import { BaseError, Options, UnknownError } from 'manekani-core';
import { Subject } from 'manekani-core';
import { TokenWithHash } from 'manekani-core';
import { PublicUser, User } from 'manekani-core';
import {
  Prisma,
  Subject as SubjectModel,
  Deck as DeckModel,
  ApiToken as TokenModel,
  User as UserModel,
} from '@prisma/client';

export const inlineAsyncTry = async <T>(
  action: () => Promise<T>,
  except: (error: unknown) => never,
): Promise<T> => {
  try {
    return await action();
  } catch (error) {
    except(error);
  }
};

export const runtimeOmit = <T extends {}, K extends keyof T>(
  data: T,
  omit: K[],
): Omit<T, K> => {
  const clone = structuredClone(data);
  for (const key of omit) {
    delete clone[key];
  }
  return clone;
};

export module PrismaErrors {
  type ErrorData = Omit<Options, 'cause'>;
  type ErrorFactory = new (errorOptions: Options) => BaseError;
  type PrismaErrorCode = `P${number}`;

  export const NotFound: PrismaErrorCode = 'P2025',
    UniqueConstraint: PrismaErrorCode = 'P2002',
    ForeignKeyConstraint: PrismaErrorCode = 'P2003',
    DependencyNotFound: PrismaErrorCode = 'P2025';

  export function match(
    error: unknown,
    options: {
      fallback: ErrorData;
      byError?: Record<PrismaErrorCode, [ErrorFactory, ErrorData]>;
    },
  ): BaseError {
    if (error instanceof BaseError) {
      return error;
    }

    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      console.debug('Prisma error was raised', {
        name: error.name,
        code: error.code,
        meta: error.meta,
        message: error.message,
      });

      if (options.byError && error.code in options.byError) {
        const exception: [ErrorFactory, ErrorData] | undefined =
          options.byError[error.code as keyof typeof options.byError];
        if (exception) {
          const [factory, errorOptions] = exception;
          return new factory({ ...errorOptions, cause: error });
        }
      }
    }

    return new UnknownError({
      cause: error,
      ...options.fallback,
    });
  }
}

type IdList = { id: string }[];

type SubjectRelations = {
  dependsOn?: IdList;
  dependedBy?: IdList;
  similarTo?: IdList;
  similarFrom?: IdList;
  deck: { ownerId: string };
};

type DeckRelations = {
  ownedBy: { id: string };
  subjects: { id: string }[];
};

type UserRelations = {
  email: string;
};

type TokenRelations = {
  owner: { id: string };
};

export module CastModel {
  export const intoUser = (model: UserModel & UserRelations): User => {
    return {
      id: model.id,
      email: model.email,
      username: model.username,
      isVerified: model.isVerified,
      createdAt: model.createdAt,
      updatedAt: model.updatedAt,
    };
  };

  export const intoPublicUser = (model: UserModel): PublicUser => {
    return {
      username: model.username,
      isVerified: model.isVerified,
    };
  };

  export const intoDeck = (model: DeckModel & DeckRelations): Deck => {
    return {
      id: model.id,
      createdAt: model.createdAt,
      updatedAt: model.updatedAt,
      name: model.name,
      description: model.description,
      ownerId: model.ownedBy.id,
      subjectIds: model.subjects.map(({ id }) => id),
    } satisfies Deck;
  };

  export const intoSubject = (function () {
    const partial = (model: SubjectModel & SubjectRelations) => ({
      id: model.id,
      createdAt: model.createdAt,
      updatedAt: model.updatedAt,
      category: model.category,
      level: model.level,
      name: model.name,
      value: model.value ? model.value : undefined,
      valueImage: model.valueImage ? model.valueImage : undefined,
      slug: model.slug,
      priority: model.priority,
      dependencies: mapIdList(model.dependsOn),
      dependents: mapIdList(model.dependedBy),
      similar: mapIdList(join(model.similarTo, model.similarFrom)),
      deckId: model.deckId,
      ownerId: model.deck.ownerId,
    });
    const c = (model: SubjectModel & SubjectRelations): Subject => {
      const base = partial(model);
      return {
        ...base,
        studyData: cast(model.studyData),
        resources: cast(model.resources),
      } satisfies Subject;
    };
    c.partial = partial;
    return c;
  })();

  export function intoToken(model: TokenModel & TokenRelations): Token {
    return {
      id: model.id,
      createdAt: model.createdAt,
      name: model.name,
      prefix: model.prefix,
      // token: model.token,
      claims: model.claims as TokenWithHash['claims'],
      isActive: model.isActive,
      usedAt: model.usedAt ? model.usedAt : undefined,
      ownerId: model.owner.id,
    };
  }

  function cast<T>(value: any): T {
    return value as T;
  }

  function mapIdList(list: IdList | null | undefined): string[] {
    return (list ?? []).map(({ id }) => id);
  }
  function join(a?: any[], b?: any[]) {
    return [...(a ?? []), ...(b ?? [])];
  }
}
