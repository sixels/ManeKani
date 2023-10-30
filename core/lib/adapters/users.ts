import { CreateUserDto, User } from '../domain/user';

import { IUsersRepository } from '../ports/users';
import { ResourceNotFoundError } from '../domain/error';
import { UserSession } from '../domain/auth';

export class UsersAdapter<R extends IUsersRepository> {
  constructor(private usersRepository: R) {}

  async getUser(session: UserSession): Promise<User> {
    const foundUser = await this.usersRepository.getUser(session);
    if (!foundUser) {
      throw new ResourceNotFoundError({
        context: { userId: session.userId },
        description: 'Could not find a user with the specified ID.',
      });
    }
    return foundUser;
  }

  createUser(user: CreateUserDto): Promise<User> {
    return this.usersRepository.createUser(user);
  }
}
