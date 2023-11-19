import { UsersAdapter } from 'manekani-core';
import { User } from 'manekani-core';
import { IUsersRepository } from 'manekani-core';
import { UsersProviderLabel } from '@/infra/database/usersDatabase.service';
import { Body, Controller, Inject, Post } from '@nestjs/common';

type UserRegisterHookBody = {
  user_id: string;
  email: string;
  created_at: string;
};

@Controller('/hooks/auth')
export class AuthController {
  private usersAdapter: UsersAdapter<any>;
  constructor(@Inject(UsersProviderLabel) usersService: IUsersRepository) {
    this.usersAdapter = new UsersAdapter(usersService);
  }

  @Post('user-register')
  async userRegister(@Body() body: UserRegisterHookBody): Promise<User> {
    console.debug('user register request', body);
    return await this.usersAdapter.createUser({
      id: body.user_id,
      email: body.email,
      createdAt: body.created_at,
    });
  }
}
