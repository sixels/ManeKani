import {
  Body,
  Controller,
  Delete,
  Get,
  HttpStatus,
  Param,
  Patch,
  Post,
  Query,
} from '@nestjs/common';

import { SubjectsService } from './subjects.service';
import {
  CreateSubjectDto,
  GetSubjectsFilters,
  PartialSubject,
  Subject,
  UpdateSubjectDto,
} from '@/core/domain/subject';
import { Authorize, UserData } from '@/api/auth/auth.decorator';
import { Response } from '@/api/response';
import { ApiTags } from '@nestjs/swagger';

@ApiTags('subjects')
@Controller()
export class SubjectsController {
  constructor(private readonly subjectsService: SubjectsService<any>) {}

  /**
   * Get all subjects
   *
   * @param filters the query filters
   * @returns a list of all subjects that matches the given filters
   */
  @Get('/v1/subjects')
  async getSubjects(
    @Query() filters: GetSubjectsFilters,
  ): Promise<Response<PartialSubject[]>> {
    return new Response(await this.subjectsService.v1GetSubjects(filters));
  }

  /**
   * Get a subject by its ID
   *
   * @param subjectId the subject's ID
   * @returns the detailed subject
   */
  @Get('/v1/subjects/:id')
  async getSubject(@Param('id') subjectId: string): Promise<Response<Subject>> {
    return new Response(await this.subjectsService.v1GetSubject(subjectId));
  }

  /**
   * Create a new subject
   *
   * @param userId the currently authenticated user's ID
   * @param deckId the deck's ID
   * @param createSubjectDto the subject to create
   * @returns the created subject with detailed information
   */
  @Post('/v1/subjects')
  @Authorize({ scopes: ['subject:create'] })
  async createSubject(
    @UserData('id') userId: string,
    @Body() createSubjectDto: CreateSubjectDto,
  ): Promise<Response<Subject>> {
    return new Response(
      await this.subjectsService.v1CreateSubject(userId, createSubjectDto),
      HttpStatus.CREATED,
    );
  }

  /**
   * Update a subject
   *
   * @param userId the currently authenticated user's ID
   * @param subjectId the subject's ID
   * @returns the updated subject with detailed information
   */
  @Patch('/v1/subjects/:id')
  @Authorize({ scopes: ['subject:update'] })
  async updateSubject(
    @UserData('id') userId: string,
    @Param('id') subjectId: string,
    @Body() updateSubjectDto: UpdateSubjectDto,
  ): Promise<Response<Subject>> {
    return new Response(
      await this.subjectsService.v1UpdateSubject(
        userId,
        subjectId,
        updateSubjectDto,
      ),
    );
  }

  /**
   * Delete a subject
   *
   * @param userId the currently authenticated user's ID
   * @param subjectId the subject's ID
   * @returns an empty api result
   */
  @Delete('/v1/subjects/:id')
  @Authorize({ scopes: ['subject:delete'] })
  async deleteSubject(
    @UserData('id') userId: string,
    @Param('id') subjectId: string,
  ): Promise<Response<null>> {
    await this.subjectsService.v1DeleteSubject(userId, subjectId);
    return new Response(null);
  }
}
