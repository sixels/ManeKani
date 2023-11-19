import {
  CreateUserDto,
  UpdateUserDto,
  CreateUserSchema,
  User,
  UsernameSchema,
  UpdateUserSchema,
} from '../domain/user';

import { IUsersRepository } from '../ports/users';
import { ResourceNotFoundError } from '../domain/error';
import { UserSession } from '../domain/auth';
import { Validator } from '../validator';

const CreateUserValidator = new Validator(CreateUserSchema);
const UpdateUserValidator = new Validator(UpdateUserSchema);
const UsernameValidator = new Validator(UsernameSchema);

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
    CreateUserValidator.validate(user);
    return this.usersRepository.createUser(user);
  }

  isUsernameAvailable(username: string): Promise<boolean> {
    UsernameValidator.validate(username);
    return this.usersRepository.isUsernameAvailable(username);
  }

  updateUser(userId: string, changes: UpdateUserDto): Promise<User> {
    UpdateUserValidator.validate(changes);

    changes.isComplete = undefined;
    const isComplete = Boolean(changes.username);

    return this.usersRepository.updateUser(
      userId,
      Object.assign(changes, { isComplete }),
    );
  }
}
