import { TokensAdapter } from 'manekani-core';
import { ITokenRespository } from 'manekani-core';
import { Inject, Injectable } from '@nestjs/common';

@Injectable()
export class TokensService<
  R extends ITokenRespository,
> extends TokensAdapter<R> {
  constructor(@Inject('TOKENS_REPOSITORY') tokensRepository: R) {
    super(tokensRepository);
  }
}
