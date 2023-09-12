import { Module } from '@nestjs/common';
import { PrismaService } from './prisma.service';
import { SubjectsDatabaseService } from './subjectsDatabase.service';
import { TokensDatabaseService } from './tokensDatabase.service';

@Module({
  providers: [PrismaService, SubjectsDatabaseService, TokensDatabaseService],
  exports: [SubjectsDatabaseService, TokensDatabaseService],
})
export class DatabaseModule {}
