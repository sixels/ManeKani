import { DynamicModule, Global, Module } from '@nestjs/common';

import { OryService } from './ory.service';
import sdk from '@ory/client';

type ModuleConfiguration = sdk.ConfigurationParameters;

@Global()
@Module({})
export class OryModule {
  static forRoot(options: ModuleConfiguration): DynamicModule {
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
