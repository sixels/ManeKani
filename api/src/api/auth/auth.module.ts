import { DatabaseModule } from '@/infra/database/database.module';
import { Module } from '@nestjs/common';
import { OryService } from '@/infra/auth/ory.service';
import { SsoAuthService } from './sso.service';
import { TokenAuthService } from './token.service';
import { TokensDatabaseService } from '@/infra/database/tokensDatabase.service';

@Module({
  imports: [DatabaseModule],
  providers: [
    { provide: 'AUTH_SSO_PROVIDER', useExisting: OryService },
    { provide: 'AUTH_TOKEN_PROVIDER', useExisting: TokensDatabaseService },
    SsoAuthService,
    TokenAuthService,
  ],
  exports: [SsoAuthService, TokenAuthService],
})
export class AuthModule {}
