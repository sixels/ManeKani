import { Injectable } from '@nestjs/common';
import { DatabaseService } from './database.service';
import { SubjectsDatabase } from 'manekani-infra-db';

export const SubjectsProviderLabel = 'SUBJECTS_REPOSITORY';

@Injectable()
export class SubjectsDatabaseService extends SubjectsDatabase {
  constructor(db: DatabaseService) {
    super(db);
  }
}
