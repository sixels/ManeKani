import { Injectable } from '@nestjs/common';
import { DatabaseService } from './database.service';
import { UsersDatabase } from 'manekani-infra-db';

export const UsersProviderLabel = 'USERS_REPOSITORY';

@Injectable()
export class UsersDatabaseService extends UsersDatabase {
  constructor(db: DatabaseService) {
    super(db);
  }
}
