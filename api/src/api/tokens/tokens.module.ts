import { AuthModule } from '../auth/auth.module';
import { DatabaseModule } from '@/infra/database/database.module';
import { Module } from '@nestjs/common';
import { TokensController } from './tokens.controller';
import { TokensDatabaseService } from '@/infra/database/tokensDatabase.service';
import { TokensService } from './tokens.service';

@Module({
  imports: [AuthModule, DatabaseModule],
  controllers: [TokensController],
  providers: [
    { provide: 'TOKENS_REPOSITORY', useExisting: TokensDatabaseService },
    TokensService,
  ],
  exports: [TokensService],
})
export class TokensModule {}
