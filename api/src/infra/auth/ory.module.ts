import { DynamicModule, Global, Module } from '@nestjs/common';

import { OryService } from './ory.service';
import { ConfigurationParameters } from 'manekani-infra-auth';

@Global()
@Module({})
export class OryModule {
  static forRoot(options: ConfigurationParameters): DynamicModule {
    return {
      module: OryModule,
      providers: [
        {
          provide: 'ORY_CLIENT_OPTIONS',
          useValue: options,
        },
        OryService,
      ],

      exports: [OryService],
    };
  }
}
