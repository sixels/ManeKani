import {
	Body,
	Controller,
	Delete,
	Get,
	Param,
	Patch,
	Post,
} from "@nestjs/common";

import { Response } from "@/api/response";
import { GenerateTokenDto, Token, UpdateTokenDto } from "@manekani/core";
import { ApiTags } from "@nestjs/swagger";
import { ApiAuthorize, UserData } from "../auth/auth.decorator";
import { TokensService } from "./tokens.service";

@ApiTags("tokens")
@Controller()
export class TokensController {
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
	constructor(private readonly subjectsService: TokensService<any>) {}

	@Get("/tokens")
	@ApiAuthorize({ loginOnly: true })
	async getTokens(
		@UserData("userId") userId: string,
	): Promise<Response<Token[]>> {
		return new Response(await this.subjectsService.getTokens(userId));
	}

	@Get("/tokens/:id")
	@ApiAuthorize({ loginOnly: true })
	async getToken(
		@UserData("userId") userId: string,
		@Param("id") tokenId: string,
	): Promise<Response<Token>> {
		return new Response(await this.subjectsService.getToken(userId, tokenId));
	}

	@Post("/tokens")
	@ApiAuthorize({ loginOnly: true })
	async generateToken(
		@UserData("userId") userId: string,
		@Body() tokenData: GenerateTokenDto,
	): Promise<Response<Token & { generatedToken: string }>> {
		return new Response(
			await this.subjectsService.generateToken(userId, tokenData),
		);
	}

	@Patch("/tokens/:id")
	@ApiAuthorize({ loginOnly: true })
	async updateToken(
		@UserData("userId") userId: string,
		@Param("id") tokenId: string,
		@Body() tokenData: UpdateTokenDto,
	): Promise<Response<Token>> {
		return new Response(
			await this.subjectsService.updateToken(userId, tokenId, tokenData),
		);
	}

	@Delete("/tokens/:id")
	@ApiAuthorize({ loginOnly: true })
	async deleteToken(
		@UserData("userId") userId: string,
		@Param("id") tokenId: string,
	): Promise<Response<void>> {
		return new Response(
			await this.subjectsService.deleteToken(userId, tokenId),
		);
	}
}
