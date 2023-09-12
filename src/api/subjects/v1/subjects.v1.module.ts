import { AuthModule } from '@/api/auth/auth.module';
import { DatabaseModule } from '@/services/database/database.module';
import { Module } from '@nestjs/common';
import { SubjectsController } from './subjects.controller';
import { SubjectsDatabaseService } from '@/services/database/subjectsDatabase.service';
import { SubjectsService } from './subjects.service';

@Module({
  imports: [AuthModule, DatabaseModule],
  controllers: [SubjectsController],
  providers: [
    { provide: 'SUBJECTS_REPOSITORY', useExisting: SubjectsDatabaseService },
    SubjectsService,
  ],
  exports: [SubjectsService],
})
export class SubjectsV1Module {}
