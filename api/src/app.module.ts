import { GraphQLJSONObject } from "graphql-type-json";

import { join } from "path";
import { ApiExceptionFilter } from "@/api/error";
import { ApolloDriver } from "@nestjs/apollo";
import { Module } from "@nestjs/common";
import { ConfigModule } from "@nestjs/config";
import { APP_FILTER } from "@nestjs/core";
import { GraphQLModule } from "@nestjs/graphql";
import { ApiModule } from "./api/api.module";
import { HooksModule } from "./hooks/hooks.module";
import { OryModule } from "./infra/auth/ory.module";

@Module({
	imports: [
		GraphQLModule.forRoot({
			driver: ApolloDriver,
			autoSchemaFile: join(process.cwd(), "src/graphql/schema.graphql"),
			resolvers: {
				JSONObject: GraphQLJSONObject,
			},

			context: ({ req }) => ({ req }),
		}),

		OryModule.forRoot(
			// TODO: get this value from config/environment
			{
				basePath: "http://kratos:4433",
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
