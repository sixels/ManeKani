import { Injectable } from '@nestjs/common';
import { DatabaseService } from './database.service';
import { TokensDatabase } from 'manekani-infra-db';

@Injectable()
export class TokensDatabaseService extends TokensDatabase {
  constructor(db: DatabaseService) {
    super(db);
  }
}
