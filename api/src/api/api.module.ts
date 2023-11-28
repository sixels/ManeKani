import { Module } from "@nestjs/common";
import { DecksV1Module } from "./decks/v1/decks.module";
import { SubjectsV1Module } from "./subjects/v1/subjects.module";
import { TokensModule } from "./tokens/tokens.module";

@Module({
	imports: [SubjectsV1Module, DecksV1Module, TokensModule],
})
export class ApiModule {}
