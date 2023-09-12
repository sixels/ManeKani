import { ApiBearerAuth, ApiCookieAuth } from '@nestjs/swagger';
import {
  AuthGuard,
  META_AUTH_METHOD_KEY,
  META_AUTH_SCOPES_KEY,
} from './auth.guard';
import {
  ExecutionContext,
  SetMetadata,
  UseGuards,
  applyDecorators,
  createParamDecorator,
} from '@nestjs/common';

import { Request } from 'express';
import { UserSession } from '@/core/domain/auth';

type AuthorizationOptions =
  | {
      /**
       * If this route is only accessible by logged in users
       */
      loginOnly?: false;
      /**
       * Defines the required scopes an api token must have to access this route
       */
      scopes?: string[];
    }
  | {
      loginOnly?: true;
      scopes?: undefined;
    };

export const Authorize = (
  { loginOnly = false, scopes = [] }: AuthorizationOptions = {
    loginOnly: false,
    scopes: [],
  },
) => {
  const method = loginOnly ? 'login' : 'any';
  const useBearerAuth = loginOnly ? ApiBearerAuth('ApiToken') : undefined;

  return applyDecorators(
    SetMetadata(META_AUTH_SCOPES_KEY, scopes ?? []),
    SetMetadata(META_AUTH_METHOD_KEY, method),
    UseGuards(AuthGuard),
    ApiCookieAuth('Login'),
    useBearerAuth ?? (() => {}),
  );
};

export const UserData = createParamDecorator(
  (data: keyof UserSession, ctx: ExecutionContext) => {
    const request: Request = ctx.switchToHttp().getRequest();
    const user: UserSession = request['user'];
    return data ? user?.[data] : user;
  },
);
