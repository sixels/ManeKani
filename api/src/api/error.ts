import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import {
  BaseError,
  ForbiddenError,
  InvalidRequestError,
  ResourceCollidesError,
  ResourceNotFoundError,
  UnauthorizedError,
} from 'manekani-core';

import { GqlExceptionFilter } from '@nestjs/graphql';
import { Response } from 'express';

export interface ApiException {
  message: string;
  error: string;
  statusCode: number;
}

@Catch(BaseError)
export class ApiExceptionFilter implements ExceptionFilter, GqlExceptionFilter {
  catch(exception: BaseError, host: ArgumentsHost) {
    const status = statusCodeFromError(exception);
    const res: ApiException = {
      message: exception.message,
      error: exception.description || 'No error information provided',
      statusCode: status,
    };

    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    if ('status' in response) {
      response.status(status).json(res);
      return;
    }

    return new HttpException(`${res.message}: ${res.error}`, res.statusCode, {
      description: res.message,
      cause: exception,
    });
  }
}

export function statusCodeFromError(error: BaseError): number {
  if (error instanceof InvalidRequestError) return HttpStatus.BAD_REQUEST;
  if (error instanceof ResourceNotFoundError) return HttpStatus.NOT_FOUND;
  if (error instanceof ResourceCollidesError) return HttpStatus.CONFLICT;
  if (error instanceof UnauthorizedError) return HttpStatus.UNAUTHORIZED;
  if (error instanceof ForbiddenError) return HttpStatus.FORBIDDEN;
  return HttpStatus.INTERNAL_SERVER_ERROR;
}
