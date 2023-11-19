import { CastModel, PrismaErrors, inlineAsyncTry } from './common';
import { CreateUserDto, PublicUser, UpdateUserDto, User } from 'manekani-core';

import { IUsersRepository } from 'manekani-core';
import { PrismaClient } from '@prisma/client';
import { UserSession } from 'manekani-core';

export const UsersProviderLabel = 'USERS_REPOSITORY';

export class UsersDatabase implements IUsersRepository {
  constructor(private client: PrismaClient) {}

  private get users() {
    return this.client.user;
  }

  async getUser(session: UserSession): Promise<User | null> {
    const foundUser = await inlineAsyncTry(
      () =>
        this.users.findUnique({
          where: { id: session.userId },
          include: {
            decks: { select: { id: true } },
          },
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { userId: session.userId },
            description: `An unknown error occurred while retrieving the user.`,
          },
        });
      },
    );

    return (
      foundUser &&
      CastModel.intoUser(Object.assign(foundUser, { email: session.email! }))
    );
  }

  async getPublicUser(username: string): Promise<PublicUser | null> {
    const foundUser = await inlineAsyncTry(
      () =>
        this.users.findUnique({
          where: { username },
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { username },
            description: 'An unknown error occurred while retrieving the user.',
          },
        });
      },
    );

    return foundUser && CastModel.intoPublicUser(foundUser);
  }

  async createUser(user: CreateUserDto): Promise<User> {
    const createdUser = await inlineAsyncTry(
      () =>
        this.users.create({
          data: {
            id: user.id,
            createdAt: user.createdAt,
            username: user.username,
          },
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { userId: user.id },
            description: `An unknown error occurred while creating the user.`,
          },
        });
      },
    );

    return CastModel.intoUser(
      Object.assign(createdUser, {
        email: user.email,
        decks: [],
      }),
    );
  }

  async isUsernameAvailable(username: string): Promise<boolean> {
    const foundUser = await inlineAsyncTry(
      () =>
        this.users.findUnique({
          where: { username },
          select: { id: true },
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { username },
            description: 'An unknown error occurred while retrieving the user.',
          },
        });
      },
    );

    return !foundUser;
  }

  async updateUser(userId: string, changes: UpdateUserDto): Promise<User> {
    const updatedUser = await inlineAsyncTry(
      () =>
        this.users.update({
          where: { id: userId },
          data: {
            username: changes.username,
            displayName: changes.displayName,
            isComplete: changes.isComplete,
          },
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { userId },
            description: 'An unknown error occurred while updating the user.',
          },
        });
      },
    );

    return CastModel.intoUser(Object.assign(updatedUser, { email: '' }));
  }
}
