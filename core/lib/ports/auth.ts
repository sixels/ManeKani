import { UserSession } from '../domain/auth';

export interface ISsoAuthenticator {
  requiredCookies(): string[];
  getCookieSession(cookies: string): Promise<UserSession>;
}
