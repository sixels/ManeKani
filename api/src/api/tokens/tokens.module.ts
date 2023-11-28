import { DatabaseModule } from "@/infra/database/database.module";
import { TokensDatabaseService } from "@/infra/database/tokensDatabase.service";
import { Module } from "@nestjs/common";
import { AuthModule } from "../auth/auth.module";
import { TokensController } from "./tokens.controller";
import { TokensService } from "./tokens.service";

@Module({
	imports: [AuthModule, DatabaseModule],
	controllers: [TokensController],
	providers: [
		{ provide: "TOKENS_REPOSITORY", useExisting: TokensDatabaseService },
		TokensService,
	],
	exports: [TokensService],
})
export class TokensModule {}
