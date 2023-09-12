import { DynamicModule, Module } from '@nestjs/common';
import { OryService } from './ory.service';
import sdk from '@ory/client';

type ModuleConfiguration = sdk.ConfigurationParameters;

@Module({})
export class OryModule {
  static register(options: ModuleConfiguration): DynamicModule {
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
