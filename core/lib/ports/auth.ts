import { UserSession } from '../domain/auth';

export interface ISsoAuthenticator {
  requiredCookies(): string[];
  getCookieSession(cookies: string): Promise<UserSession>;

  registerUsername(userId: string, username: string): Promise<void>;
  updateUsername(userId: string, username: string): Promise<void>;
}
