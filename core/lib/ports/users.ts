import { CreateUserDto, PublicUser, UpdateUserDto, User } from '../domain/user';

import { UserSession } from '../domain/auth';

export interface IUsersRepository {
  getUser(session: UserSession): Promise<User | null>;
  getPublicUser(username: string): Promise<PublicUser | null>;
  createUser(user: CreateUserDto): Promise<User>;
  updateUser(userId: string, changes: UpdateUserDto): Promise<User>;

  isUsernameAvailable(username: string): Promise<boolean>;
}
