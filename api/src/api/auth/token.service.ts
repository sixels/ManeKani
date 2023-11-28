import { ITokenRespository, TokenAuthAdapter } from "@manekani/core";
import { Inject, Injectable } from "@nestjs/common";

@Injectable()
export class TokenAuthService<
	R extends ITokenRespository,
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
> extends TokenAuthAdapter<any> {
	constructor(@Inject("AUTH_TOKEN_PROVIDER") provider: R) {
		super(provider);
	}
}
