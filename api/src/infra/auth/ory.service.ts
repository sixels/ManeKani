import { ISsoAuthenticator } from "@manekani/core";
import {
	ConfigurationParameters,
	SsoAuthenticator,
} from "@manekani/infra-auth";
import { Inject, Injectable } from "@nestjs/common";

export const ORY_OPTIONS_KEY = "ORY_CLIENT_OPTIONS";

@Injectable()
export class OryService extends SsoAuthenticator implements ISsoAuthenticator {
	constructor(@Inject(ORY_OPTIONS_KEY) options: ConfigurationParameters) {
		super(options);
	}
}
