import Ajv, { ValidateFunction } from 'ajv';
import { Static, TSchema, Type } from '@sinclair/typebox';

import { InvalidRequestError } from '@/core/domain/error';
import ajvExtErrors from 'ajv-errors';
import ajvExtFormats from 'ajv-formats';

const FAIL_MESSAGE = 'Schema validation failed';

export class Validator<S extends TSchema, T = Static<S>> {
  private readonly check: ValidateFunction<S>;

  constructor(schema: S) {
    const ajv = ajvExtFormats(
      ajvExtErrors(new Ajv({ useDefaults: true, allErrors: true })),
    );
    this.check = ajv.compile(schema);
  }

  isValid(data: unknown): data is T {
    return this.check(data);
  }

  validate(data: unknown): T {
    if (this.isValid(data)) {
      return data;
    }

    const errors = this.check.errors;
    const errorsString =
      errors?.map(({ message }) => message).join('; ') ?? FAIL_MESSAGE;

    throw new InvalidRequestError({
      cause: new Error(FAIL_MESSAGE),
      context: { errors },
      description: errorsString,
    });
  }
}

const UuidSchema = Type.String({ format: 'uuid' });
export const UuidValidator = new Validator(UuidSchema);
