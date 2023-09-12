import {
  BaseError,
  InvalidRequestError,
  ResourceCollidesError,
  ResourceNotFoundError,
  UnauthorizedError,
  UnknownError,
} from '@/core/domain/error';
import { HttpStatus } from '@nestjs/common';
import { statusCodeFromError } from './error';

export const TypedCoreExceptions = (exceptions: {
  [name in string]: BaseError;
}) => {
  const exceptionsToApply = Object.entries(exceptions).map(
    ([name, exception]) => {
      const statusCode = statusCodeFromError(exception);
      //   return TypedException<${name}>(HttpStatus.${statusCode}, '${exception.message}')
    },
  );
  //   for (const exception of exceptions) {
  //   }
};
