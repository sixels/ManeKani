import * as cookieParser from 'cookie-parser';
import * as request from 'supertest';

import { createMock } from '@golevelup/ts-jest';
import { INestApplication } from '@nestjs/common';
import { APP_FILTER } from '@nestjs/core';
import { Test } from '@nestjs/testing';

import { ApiExceptionFilter } from '@/api/error';
import { SubjectsV1Module } from '@/api/subjects/v1/subjects.v1.module';
import {
  ResourceCollidesError,
  ResourceNotFoundError,
  UnknownError,
} from '@/core/domain/error';
import {
  CreateSubjectDto,
  Subject,
  UpdateSubjectDto,
} from '@/core/domain/subject';
import { ISubjectRepositoryV1 } from '@/core/ports/subjects';
import { PrismaService } from '@/services/database/prisma.service';
import { SubjectsDatabaseService } from '@/services/database/subjectsDatabase.service';
import { TokensDatabaseService } from '@/services/database/tokensDatabase.service';
import { OryService } from '@/services/ory/ory.service';
import {
  EMPTY_UUID,
  TEST_API_TOKEN,
  TEST_UUID,
  UNPRIVILEGED_API_TOKEN,
  mockSsoService,
  mockTokenService,
} from './common';

describe('Subjects', () => {
  let app: INestApplication;
  const repositoryMock = createMock<ISubjectRepositoryV1>();

  const makeRequest = () => request(app.getHttpServer());

  const testSubject: Subject = {
    ...createMock<Subject>(),
    id: TEST_UUID,
    level: 1,
    name: 'test',
    value: 'test',
    slug: 'test',
    category: 'test',
    priority: 1,
    deckId: TEST_UUID,
  };

  beforeAll(async () => {
    const moduleRef = await Test.createTestingModule({
      imports: [SubjectsV1Module],
      providers: [{ provide: APP_FILTER, useClass: ApiExceptionFilter }],
    })
      .overrideProvider(PrismaService)
      .useValue(createMock())
      .overrideProvider(SubjectsDatabaseService)
      .useValue(repositoryMock)
      .overrideProvider(TokensDatabaseService)
      .useValue(await mockTokenService())
      .overrideProvider(OryService)
      .useValue(mockSsoService())
      .compile();

    app = moduleRef.createNestApplication();
    app.use(cookieParser());
    await app.init();
  });

  describe('GET /subjects', () => {});

  describe('GET /subjects/:id', () => {
    repositoryMock.v1GetSubject.mockImplementation(async (subjectId: string) =>
      subjectId == testSubject.id ? testSubject : null,
    );

    it('should return a subject by its id', () => {
      return makeRequest()
        .get(`/v1/subjects/${testSubject.id}`)
        .expect(200)
        .expect({ data: testSubject, statusCode: 200 });
    });
    it("should return a 404 if there's not subject with the provided id", () => {
      return makeRequest().get(`/v1/subjects/${EMPTY_UUID}`).expect(404);
    });
    it("should return a 400 if the provided id isn't a valid uuid", () => {
      return makeRequest().get(`/v1/subjects/invalid-uuid`).expect(400);
    });
    it('should return a 500 if the repository throws an unknown error', () => {
      repositoryMock.v1GetSubject.mockImplementationOnce(async (_: string) => {
        throw new UnknownError({});
      });

      return makeRequest().get(`/v1/subjects/${testSubject.id}`).expect(500);
    });
  });

  describe('POST /subjects', () => {
    const testCreateSubject = Object.freeze({
      ...createMock<CreateSubjectDto>(),
      level: 1,
      name: 'test',
      value: 'test',
      slug: 'test',
      category: 'test',
      priority: 1,
      deckId: EMPTY_UUID,
    } satisfies CreateSubjectDto);

    repositoryMock.v1CreateSubject.mockImplementation(
      async (_: string, subject: CreateSubjectDto) => ({
        ...createMock<Subject>(),
        ...structuredClone(subject),
      }),
    );

    it('should create a subject', () => {
      return makeRequest()
        .post(`/v1/subjects`)
        .set('Cookie', 'foo=bar')
        .send(testCreateSubject)
        .expect(201)
        .expect({
          data: {
            ...testCreateSubject,
            studyData: [],
            additionalStudyData: {},
            resources: [],
            dependents: [],
            dependencies: [],
            similar: [],
          },
          statusCode: 201,
        });
    });
    it('should create a subject with api token', () => {
      return makeRequest()
        .post(`/v1/subjects`)
        .set('Authorization', `Bearer ${TEST_API_TOKEN}`)
        .send(testCreateSubject)
        .expect(201)
        .expect({
          data: {
            ...testCreateSubject,
            studyData: [],
            additionalStudyData: {},
            resources: [],
            dependents: [],
            dependencies: [],
            similar: [],
          },
          statusCode: 201,
        });
    });
    it('should return a 400 if the provided subject is invalid', () => {
      const testCreateSubjectInvalid = Object.freeze({
        ...testCreateSubject,
        level: 0,
        slug: '@invalid#slug',
        priority: -1,
        deckId: 'a',
      } satisfies CreateSubjectDto);

      return makeRequest()
        .post(`/v1/subjects`)
        .set('Cookie', 'foo=bar')
        .send(testCreateSubjectInvalid)
        .expect(400);
    });
    it("should return a 401 if the user isn't authenticated", () => {
      // sending request without cookies
      return makeRequest()
        .post(`/v1/subjects`)
        .send(testCreateSubject)
        .expect(401);
    });
    it('should return a 403 if the user has no privilege to create a subject', () => {
      return makeRequest()
        .post(`/v1/subjects`)
        .set('Authorization', `Bearer ${UNPRIVILEGED_API_TOKEN}`)
        .send(testCreateSubject)
        .expect(403);
    });
    it('should return a 409 if the subject already exists', () => {
      repositoryMock.v1CreateSubject.mockImplementationOnce(
        async (_: string, _subject: Subject) => {
          throw new ResourceCollidesError({});
        },
      );

      return makeRequest()
        .post(`/v1/subjects`)
        .set('Cookie', 'foo=bar')
        .send(testCreateSubject)
        .expect(409);
    });
    it('should return a 500 if the repository throws an error', () => {
      repositoryMock.v1CreateSubject.mockImplementationOnce(
        async (_: string, _subject: Subject) => {
          throw new UnknownError({});
        },
      );

      return makeRequest()
        .post(`/v1/subjects`)
        .set('Cookie', 'foo=bar')
        .send(testCreateSubject)
        .expect(500);
    });
  });

  describe('PATCH /subjects/:id', () => {
    const testUpdateSubject: UpdateSubjectDto = {
      level: 2,
      name: 'test patched',
      priority: 1,
    };

    repositoryMock.v1UpdateSubject.mockImplementation(
      async (_: string, subjectId: string, changes: UpdateSubjectDto) => {
        if (subjectId == testSubject.id) {
          return {
            ...testSubject,
            ...changes,
          } satisfies Subject;
        }
        throw new ResourceNotFoundError({});
      },
    );

    it('should update a subject', () => {
      return makeRequest()
        .patch(`/v1/subjects/${testSubject.id}`)
        .set('Cookie', 'foo=bar')
        .send(testUpdateSubject)
        .expect(200)
        .expect({
          data: {
            ...testSubject,
            ...testUpdateSubject,
            studyData: [],
            additionalStudyData: {},
            resources: [],
            dependents: [],
            dependencies: [],
            similar: [],
          },
          statusCode: 200,
        });
    });
    it('should update a subject with api token', () => {
      return makeRequest()
        .patch(`/v1/subjects/${testSubject.id}`)
        .set('Authorization', `Bearer ${TEST_API_TOKEN}`)
        .send(testUpdateSubject)
        .expect(200);
    });
    it("should return a 400 if the provided id isn't a valid uuid", () => {
      return makeRequest()
        .patch(`/v1/subjects/invalid-uuid`)
        .set('Cookie', 'foo=bar')
        .send(testUpdateSubject)
        .expect(400);
    });
    it('should return a 400 if the provided subject is invalid', () => {
      const testUpdateSubjectInvalid: UpdateSubjectDto = {
        level: 0,
        slug: '@invalid#slug',
        priority: -1,
      };

      return makeRequest()
        .patch(`/v1/subjects/${testSubject.id}`)
        .set('Cookie', 'foo=bar')
        .send(testUpdateSubjectInvalid)
        .expect(400);
    });
    it("should return a 401 if the user isn't authenticated", () => {
      return makeRequest()
        .patch(`/v1/subjects/${testSubject.id}`)
        .send(testUpdateSubject)
        .expect(401);
    });
    it('should return a 403 if the user has no privilege to update a subject', () => {
      return makeRequest()
        .patch(`/v1/subjects/${testSubject.id}`)
        .set('Authorization', `Bearer ${UNPRIVILEGED_API_TOKEN}`)
        .send(testUpdateSubject)
        .expect(403);
    });
    it("should return a 404 if the subject doesn't exist", () => {
      return makeRequest()
        .patch(`/v1/subjects/${EMPTY_UUID}`)
        .set('Cookie', 'foo=bar')
        .send(testUpdateSubject)
        .expect(404);
    });
    it('should return a 500 if the repository throws an error', () => {
      repositoryMock.v1UpdateSubject.mockImplementationOnce(
        async (_uid: string, _id: string, _subject: Subject) => {
          throw new UnknownError({});
        },
      );

      return makeRequest()
        .patch(`/v1/subjects/${testSubject.id}`)
        .set('Cookie', 'foo=bar')
        .send(testUpdateSubject)
        .expect(500);
    });
  });

  describe('DELETE /subjects/:id', () => {
    repositoryMock.v1DeleteSubject.mockImplementation(
      async (_: string, subjectId: string) => {
        if (subjectId == TEST_UUID) {
          return;
        }
        throw new ResourceNotFoundError({});
      },
    );

    it('should delete a subject by its id', () => {
      return makeRequest()
        .delete(`/v1/subjects/${TEST_UUID}`)
        .set('Cookie', 'foo=bar')
        .expect(200)
        .expect({
          data: null,
          statusCode: 200,
        });
    });
    it('should delete a subject with api token', () => {
      return makeRequest()
        .delete(`/v1/subjects/${testSubject.id}`)
        .set('Authorization', `Bearer ${TEST_API_TOKEN}`)
        .expect(200);
    });
    it("should return a 400 if the provided id isn't a valid uuid", () => {
      return makeRequest()
        .delete(`/v1/subjects/invalid-uuid`)
        .set('Cookie', 'foo=bar')
        .expect(400);
    });
    it("should return a 401 if the user isn't authenticated", () => {
      return makeRequest().delete(`/v1/subjects/${testSubject.id}`).expect(401);
    });
    it('should return a 403 if the user has no privilege to delete a subject', () => {
      return makeRequest()
        .delete(`/v1/subjects/${testSubject.id}`)
        .set('Authorization', `Bearer ${UNPRIVILEGED_API_TOKEN}`)
        .expect(403);
    });
    it("should return a 404 if the subject doesn't exist", () => {
      return makeRequest()
        .delete(`/v1/subjects/${EMPTY_UUID}`)
        .set('Cookie', 'foo=bar')
        .expect(404);
    });
    it('should return a 500 if the repository throws an error', () => {
      repositoryMock.v1DeleteSubject.mockImplementationOnce(
        async (_: string) => {
          throw new UnknownError({});
        },
      );

      return makeRequest()
        .delete(`/v1/subjects/${testSubject.id}`)
        .set('Cookie', 'foo=bar')
        .expect(500);
    });
  });

  afterAll(async () => {
    await app.close();
  });
});
