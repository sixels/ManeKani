export type Options = {
  cause?: unknown;
  description?: string;
  readonly context?: unknown;
};

export class BaseError extends Error {
  public cause?: Options['cause'];
  public description?: Options['description'];
  public readonly context?: Options['context'];

  constructor(
    public message: string,
    { cause, description, context }: Options = {},
  ) {
    super(message);
    this.cause = cause;
    this.description = description;
    this.context = context;
  }
}

export class InvalidRequestError extends BaseError {
  constructor(options: Options) {
    super('Invalid request', options);
  }
}
export class ResourceNotFoundError extends BaseError {
  constructor(options: Options) {
    super('Resource not found', options);
  }
}
export class ResourceCollidesError extends BaseError {
  constructor(options: Options) {
    super('Resource already exists', options);
  }
}
export class UnauthorizedError extends BaseError {
  constructor(options: Options) {
    super('Unauthorized request', options);
  }
}
export class ForbiddenError extends BaseError {
  constructor(options: Options) {
    super('Forbidden request', options);
  }
}
export class UnknownError extends BaseError {
  constructor(options: Options) {
    super('Unknown error occurred', options);
  }
}
