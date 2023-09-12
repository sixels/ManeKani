import { ISsoAuthenticator } from '@/core/ports/auth';
import { SsoAuthAdapter } from '@/core/adapters/auth';
import { Inject, Injectable } from '@nestjs/common';

@Injectable()
export class SsoAuthService<
  R extends ISsoAuthenticator,
> extends SsoAuthAdapter<any> {
  constructor(@Inject('AUTH_SSO_PROVIDER') provider: R) {
    super(provider);
  }
}
