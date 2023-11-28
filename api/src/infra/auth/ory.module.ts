import { DynamicModule, Global, Module } from "@nestjs/common";

import { ConfigurationParameters } from "@manekani/infra-auth";
import { OryService } from "./ory.service";

@Global()
@Module({})
// biome-ignore lint/complexity/noStaticOnlyClass: This is required by Nest
export class OryModule {
	static forRoot(options: ConfigurationParameters): DynamicModule {
		return {
			module: OryModule,
			providers: [
				{
					provide: "ORY_CLIENT_OPTIONS",
					useValue: options,
				},
				OryService,
			],

			exports: [OryService],
		};
	}
}
