import * as cookieParser from "cookie-parser";

import { NestFactory } from "@nestjs/core";
import { NestExpressApplication } from "@nestjs/platform-express";
import { AppModule } from "./app.module";

async function bootstrap() {
	const app = await NestFactory.create<NestExpressApplication>(AppModule);

	app.use(
		// we ain't signing any cookies
		cookieParser(),
	);

	await app.listen(3000);
}
bootstrap();
