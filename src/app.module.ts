import { APP_FILTER } from '@nestjs/core';
import { ApiExceptionFilter } from '@/api/error';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule } from '@nestjs/config';
import { Module } from '@nestjs/common';
import { SubjectsV1Module } from '@/api/subjects/v1/subjects.v1.module';
import { TokensModule } from './api/tokens/tokens.module';

@Module({
  imports: [
    SubjectsV1Module,
    TokensModule,
    ConfigModule.forRoot({
      envFilePath: `.env.${process.env.NODE_ENV}`,
    }),
  ],
  controllers: [AppController],
  providers: [
    { provide: APP_FILTER, useClass: ApiExceptionFilter },
    AppService,
  ],
})
export class AppModule {}
