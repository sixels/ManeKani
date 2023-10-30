import { CastModel, PrismaErrors, inlineAsyncTry } from './common';
import { CreateTokenDto, Token, UpdateTokenDto } from 'manekani-core';
import { ResourceCollidesError, ResourceNotFoundError } from 'manekani-core';

import { ITokenRespository } from 'manekani-core';
import { PrismaClient } from '@prisma/client';

export class TokensDatabase implements ITokenRespository {
  constructor(private client: PrismaClient) {}
  private get tokens() {
    return this.client.apiToken;
  }

  async getTokens(userId: string): Promise<Token[]> {
    const tokens = await inlineAsyncTry(
      () =>
        this.tokens.findMany({
          where: { ownerId: userId },
          include: includeTokenOwner,
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { userId },
            description: `An unknown error occurred while retrieving the user's tokens.`,
          },
          byError: {
            [PrismaErrors.NotFound]: [
              ResourceNotFoundError,
              {
                context: { userId },
                description: `No tokens were found.`,
              },
            ],
          },
        });
      },
    );

    return tokens.map(CastModel.intoToken);
  }

  async getToken(userId: string, tokenId: string): Promise<Token | null> {
    const token = await inlineAsyncTry(
      () =>
        this.tokens.findUnique({
          where: { id: tokenId, ownerId: userId },
          include: includeTokenOwner,
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { tokenId },
            description: `An unknown error occurred while retrieving the token with id "${tokenId}".`,
          },
          byError: {
            [PrismaErrors.NotFound]: [
              ResourceNotFoundError,
              {
                context: { tokenId },
                description: `No Tokens with id "${tokenId}" were found.`,
              },
            ],
          },
        });
      },
    );

    return token && CastModel.intoToken(token);
  }

  async createToken(userId: string, data: CreateTokenDto): Promise<Token> {
    const token = await inlineAsyncTry(
      () =>
        this.tokens.create({
          data: {
            ...data,
            ownerId: userId,
          },
          include: includeTokenOwner,
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { data },
            description: `An unknown error occurred while creating the token.`,
          },
          byError: {
            [PrismaErrors.UniqueConstraint]: [
              ResourceCollidesError,
              {
                context: { data },
                description: `You already have a token with name "${data.name}".`,
              },
            ],
          },
        });
      },
    );

    return CastModel.intoToken(token);
  }
  async updateToken(
    userId: string,
    tokenId: string,
    changes: UpdateTokenDto,
  ): Promise<Token> {
    const token = await inlineAsyncTry(
      () =>
        this.tokens.update({
          where: { id: tokenId, ownerId: userId },
          data: {
            ...changes,
          },
          include: includeTokenOwner,
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { tokenId, changes },
            description: `An unknown error occurred while updating the token with id "${tokenId}".`,
          },
          byError: {
            [PrismaErrors.NotFound]: [
              ResourceNotFoundError,
              {
                context: { tokenId },
                description: `No tokens with id "${tokenId}" were found.`,
              },
            ],
            [PrismaErrors.UniqueConstraint]: [
              ResourceCollidesError,
              {
                context: { changes },
                description: `You already have a token with name "${changes.name}".`,
              },
            ],
          },
        });
      },
    );

    return CastModel.intoToken(token);
  }
  async deleteToken(userId: string, id: string): Promise<void> {
    await inlineAsyncTry(
      () =>
        this.tokens.delete({
          where: { id: id, ownerId: userId },
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { id },
            description: `An unknown error occurred while deleting the token with id "${id}".`,
          },
          byError: {
            [PrismaErrors.NotFound]: [
              ResourceNotFoundError,
              {
                context: { id },
                description: `No tokens with id "${id}" were found.`,
              },
            ],
          },
        });
      },
    );
  }

  async useToken(token: string): Promise<Token> {
    const foundToken = await inlineAsyncTry(
      () =>
        this.tokens.update({
          where: { token },
          data: { usedAt: new Date() },
          include: includeTokenOwner,
        }),
      (error) => {
        throw PrismaErrors.match(error, {
          fallback: {
            context: { token },
            description: `An unknown error occurred while retrieving the token with id "${token}".`,
          },
          byError: {
            [PrismaErrors.NotFound]: [
              ResourceNotFoundError,
              {
                context: { token },
                description: `No Tokens with id "${token}" were found.`,
              },
            ],
          },
        });
      },
    );

    return CastModel.intoToken(foundToken);
  }
}

const includeTokenOwner = { owner: { select: { id: true } } };
