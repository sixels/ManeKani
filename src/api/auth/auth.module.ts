import { SsoAuthAdapter, TokenAuthAdapter } from '@/core/adapters/auth';

import { DatabaseModule } from '@/services/database/database.module';
import { Module } from '@nestjs/common';
import { OryModule } from '@/services/ory/ory.module';
import { OryService } from '@/services/ory/ory.service';
import { SsoAuthService } from './sso.service';
import { TokenAuthService } from './token.service';
import { TokensDatabaseService } from '@/services/database/tokensDatabase.service';

@Module({
  imports: [
    OryModule.register(
      // TODO: get this value from config/environment
      {
        basePath: 'http://kratos:4433',
        baseOptions: {
          withCredentials: true,
        },
      },
    ),
    DatabaseModule,
  ],
  providers: [
    { provide: 'AUTH_SSO_PROVIDER', useExisting: OryService },
    { provide: 'AUTH_TOKEN_PROVIDER', useExisting: TokensDatabaseService },
    SsoAuthService,
    TokenAuthService,
  ],
  exports: [SsoAuthService, TokenAuthService],
})
export class AuthModule {}
