import { CreateTokenDto, TokenWithHash, UpdateTokenDto } from '../domain/token';

export interface ITokenRespository {
  // todo: implement filters
  getTokens(userId: string): Promise<TokenWithHash[]>;
  getToken(userId: string, tokenId: string): Promise<TokenWithHash | null>; // requires user id for extra security
  createToken(userId: string, data: CreateTokenDto): Promise<TokenWithHash>;
  updateToken(
    userId: string,
    tokenId: string,
    changes: UpdateTokenDto,
  ): Promise<TokenWithHash>;
  deleteToken(userId: string, tokenId: string): Promise<void>;
}
