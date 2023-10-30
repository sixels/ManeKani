import { ITokenRespository } from 'manekani-core';
import { TokenAuthAdapter } from 'manekani-core';
import { Inject, Injectable } from '@nestjs/common';

@Injectable()
export class TokenAuthService<
  R extends ITokenRespository,
> extends TokenAuthAdapter<any> {
  constructor(@Inject('AUTH_TOKEN_PROVIDER') provider: R) {
    super(provider);
  }
}
