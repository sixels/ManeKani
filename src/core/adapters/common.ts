import { InvalidRequestError } from '../domain/error';
import { UuidValidator } from '@/lib/validator';

export function validateId(id: string) {
  if (!UuidValidator.isValid(id)) {
    throw new InvalidRequestError({
      cause: new Error('UUID validation failed'),
      context: { id },
      description: 'The provided ID is not a valid UUID',
    });
  }
}
