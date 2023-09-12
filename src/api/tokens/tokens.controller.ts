import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  Patch,
  Post,
} from '@nestjs/common';

import { Authorize, UserData } from '../auth/auth.decorator';
import { GenerateTokenDto, Token, UpdateTokenDto } from '@/core/domain/token';
import { TokensService } from './tokens.service';
import { Response } from '@/api/response';
import { ApiTags } from '@nestjs/swagger';

@ApiTags('tokens')
@Controller()
export class TokensController {
  constructor(private readonly subjectsService: TokensService<any>) {}

  @Get('/tokens')
  @Authorize({ loginOnly: true })
  async getTokens(@UserData('id') userId: string): Promise<Response<Token[]>> {
    return new Response(await this.subjectsService.getTokens(userId));
  }

  @Get('/tokens/:id')
  @Authorize({ loginOnly: true })
  async getToken(
    @UserData('id') userId: string,
    @Param('id') tokenId: string,
  ): Promise<Response<Token>> {
    return new Response(await this.subjectsService.getToken(userId, tokenId));
  }

  @Post('/tokens')
  @Authorize({ loginOnly: true })
  async generateToken(
    @UserData('id') userId: string,
    @Body() tokenData: GenerateTokenDto,
  ): Promise<Response<Token & { generatedToken: string }>> {
    return new Response(
      await this.subjectsService.generateToken(userId, tokenData),
    );
  }

  @Patch('/tokens/:id')
  @Authorize({ loginOnly: true })
  async updateToken(
    @UserData('id') userId: string,
    @Param('id') tokenId: string,
    @Body() tokenData: UpdateTokenDto,
  ): Promise<Response<Token>> {
    return new Response(
      await this.subjectsService.updateToken(userId, tokenId, tokenData),
    );
  }

  @Delete('/tokens/:id')
  @Authorize({ loginOnly: true })
  async deleteToken(
    @UserData('id') userId: string,
    @Param('id') tokenId: string,
  ): Promise<Response<void>> {
    return new Response(
      await this.subjectsService.deleteToken(userId, tokenId),
    );
  }
}
