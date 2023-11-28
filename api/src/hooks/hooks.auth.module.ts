import {
	UsersDatabaseService,
	UsersProviderLabel,
} from "@/infra/database/usersDatabase.service";

import { DatabaseModule } from "@/infra/database/database.module";
import { Module } from "@nestjs/common";
import { AuthController } from "./hooks.auth.controller";

@Module({
	imports: [DatabaseModule],
	controllers: [AuthController],
	providers: [
		{ provide: UsersProviderLabel, useExisting: UsersDatabaseService },
	],
	exports: [],
})
export class AuthModule {}
