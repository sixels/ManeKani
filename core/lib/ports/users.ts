import { CreateUserDto, PublicUser, User } from '../domain/user';

import { UserSession } from '../domain/auth';

export interface IUsersRepository {
  getUser(session: UserSession): Promise<User | null>;
  getPublicUser(username: string): Promise<PublicUser | null>;

  createUser(user: CreateUserDto): Promise<User>;
}
