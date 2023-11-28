import { CreateTokenDto, Token, UpdateTokenDto } from "../domain/token";

export interface ITokenRespository {
	// todo: implement filters
	getTokens(userId: string): Promise<Token[]>;
	getToken(userId: string, tokenId: string): Promise<Token | null>;
	createToken(userId: string, data: CreateTokenDto): Promise<Token>;
	updateToken(
		userId: string,
		tokenId: string,
		changes: UpdateTokenDto,
	): Promise<Token>;
	deleteToken(userId: string, tokenId: string): Promise<void>;
	useToken(token: string): Promise<Token>;
}
