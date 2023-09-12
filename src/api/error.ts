import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpStatus,
} from '@nestjs/common';
import {
  BaseError,
  ForbiddenError,
  InvalidRequestError,
  ResourceCollidesError,
  ResourceNotFoundError,
  UnauthorizedError,
} from '@/core/domain/error';

import { Response } from 'express';

export interface ApiException {
  message: string;
  error: string;
  statusCode: number;
}

@Catch(BaseError)
export class ApiExceptionFilter implements ExceptionFilter {
  catch(exception: BaseError, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();

    const status = statusCodeFromError(exception);
    const res: ApiException = {
      message: exception.message,
      error: exception.description || 'No error information provided',
      statusCode: status,
    };

    const request = ctx.getRequest<Request>();
    response.status(status).json(res);
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
