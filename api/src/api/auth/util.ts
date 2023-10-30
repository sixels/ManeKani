import { ExecutionContext } from '@nestjs/common';
import { GqlExecutionContext } from '@nestjs/graphql';
import { Request } from 'express';
import { UserSession } from 'manekani-core';

type RequestWithUser = Request & { user?: UserSession };

export function getRequest(context: ExecutionContext): RequestWithUser {
  if ((context.getType() as string) == 'graphql') {
    const ctx = GqlExecutionContext.create(context);
    return ctx.getContext().req;
  } else {
    return context.switchToHttp().getRequest();
  }
}
