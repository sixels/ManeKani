import { Token } from '../domain/token';
import { UserSession } from '../domain/auth';

export interface ISsoAuthenticator {
  requiredCookies(): string[];
  getSessionFromCookies(cookies: string): Promise<UserSession>;
}

export interface ITokenAuthenticator {
  useToken(token: string): Promise<Token>;
}
