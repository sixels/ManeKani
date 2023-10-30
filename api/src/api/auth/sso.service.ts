import { ISsoAuthenticator } from 'manekani-core';
import { SsoAuthAdapter } from 'manekani-core';
import { Inject, Injectable } from '@nestjs/common';

@Injectable()
export class SsoAuthService<
  R extends ISsoAuthenticator,
> extends SsoAuthAdapter<any> {
  constructor(@Inject('AUTH_SSO_PROVIDER') provider: R) {
    super(provider);
  }
}
