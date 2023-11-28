import { UsersProviderLabel } from "@/infra/database/usersDatabase.service";
import { IUsersRepository, User, UsersAdapter } from "@manekani/core";
import { Body, Controller, Inject, Post } from "@nestjs/common";

type UserRegisterHookBody = {
	user_id: string;
	email: string;
	created_at: string;
};

@Controller("/hooks/auth")
export class AuthController {
	// biome-ignore lint/suspicious/noExplicitAny: any is just a placeholder type
	private usersAdapter: UsersAdapter<any>;
	constructor(@Inject(UsersProviderLabel) usersService: IUsersRepository) {
		this.usersAdapter = new UsersAdapter(usersService);
	}

	@Post("user-register")
	async userRegister(@Body() body: UserRegisterHookBody): Promise<User> {
		console.debug("user register request", body);
		return await this.usersAdapter.createUser({
			id: body.user_id,
			email: body.email,
			createdAt: body.created_at,
		});
	}
}
