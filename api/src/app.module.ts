import { GraphQLJSONObject } from 'graphql-type-json';

import { APP_FILTER } from '@nestjs/core';
import { ApiExceptionFilter } from '@/api/error';
import { ApolloDriver } from '@nestjs/apollo';
import { ConfigModule } from '@nestjs/config';
import { GraphQLModule } from '@nestjs/graphql';
import { HooksModule } from './hooks/hooks.module';
import { Module } from '@nestjs/common';
import { OryModule } from './infra/auth/ory.module';
import { join } from 'path';
import { ApiModule } from './api/api.module';

@Module({
  imports: [
    GraphQLModule.forRoot({
      driver: ApolloDriver,
      autoSchemaFile: join(process.cwd(), 'src/graphql/schema.gql'),
      resolvers: {
        JSONObject: GraphQLJSONObject,
      },

      context: ({ req }) => ({ req }),
    }),

    OryModule.forRoot(
      // TODO: get this value from config/environment
      {
        basePath: 'http://kratos:4433',
        baseOptions: {
          withCredentials: true,
        },
      },
    ),

    ConfigModule.forRoot({
      envFilePath: `.env.${process.env.NODE_ENV}`,
    }),

    HooksModule,
    ApiModule,
  ],
  providers: [{ provide: APP_FILTER, useClass: ApiExceptionFilter }],
})
export class AppModule {}
