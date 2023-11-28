import { ISsoAuthenticator, SsoAuthAdapter } from "@manekani/core";
import { Inject, Injectable } from "@nestjs/common";

@Injectable()
export class SsoAuthService<
	R extends ISsoAuthenticator,
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
> extends SsoAuthAdapter<any> {
	constructor(@Inject("AUTH_SSO_PROVIDER") provider: R) {
		super(provider);
	}
}
