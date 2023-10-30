import { DecksDatabase } from 'manekani-infra-db';
import { Injectable } from '@nestjs/common';
import { DatabaseService } from './database.service';

export const DecksProviderLabel = 'DECKS_REPOSITORY';

@Injectable()
export class DecksDatabaseService extends DecksDatabase {
  constructor(db: DatabaseService) {
    super(db);
  }
}
