import { ITokenAuthenticator } from '@/core/ports/auth';
import { TokenAuthAdapter } from '@/core/adapters/auth';
import { Inject, Injectable } from '@nestjs/common';

@Injectable()
export class TokenAuthService<
  R extends ITokenAuthenticator,
> extends TokenAuthAdapter<any> {
  constructor(@Inject('AUTH_TOKEN_PROVIDER') provider: R) {
    super(provider);
  }
}
