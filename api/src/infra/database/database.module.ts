import { DecksDatabaseService } from './decksDatabase.service';
import { Module } from '@nestjs/common';
import { DatabaseService } from './database.service';
import { SubjectsDatabaseService } from './subjectsDatabase.service';
import { TokensDatabaseService } from './tokensDatabase.service';
import { UsersDatabaseService } from './usersDatabase.service';

@Module({
  providers: [
    DatabaseService,
    SubjectsDatabaseService,
    DecksDatabaseService,
    TokensDatabaseService,
    UsersDatabaseService,
  ],
  exports: [
    SubjectsDatabaseService,
    DecksDatabaseService,
    UsersDatabaseService,
    TokensDatabaseService,
  ],
})
export class DatabaseModule {}
