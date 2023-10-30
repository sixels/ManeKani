import {
  DecksDatabaseService,
  DecksProviderLabel,
} from '@/infra/database/decksDatabase.service';
import {
  SubjectsDatabaseService,
  SubjectsProviderLabel,
} from '@/infra/database/subjectsDatabase.service';

import { AuthModule } from '@/api/auth/auth.module';
import { DatabaseModule } from '@/infra/database/database.module';
import { Module } from '@nestjs/common';
import { SubjectsResolver } from './subjects.resolvers';

@Module({
  imports: [AuthModule, DatabaseModule],
  controllers: [],
  providers: [
    { provide: SubjectsProviderLabel, useExisting: SubjectsDatabaseService },
    { provide: DecksProviderLabel, useExisting: DecksDatabaseService },
    SubjectsResolver,
  ],
  exports: [SubjectsResolver],
})
export class SubjectsV1Module {}
