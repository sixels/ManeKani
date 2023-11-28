import { Module } from "@nestjs/common";
import { AuthModule } from "./hooks.auth.module";

@Module({
	imports: [AuthModule],
	controllers: [],
	providers: [],
	exports: [],
})
export class HooksModule {}
