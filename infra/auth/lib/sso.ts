import { UserSession } from 'manekani-core';
import { UnauthorizedError } from 'manekani-core';
import { ISsoAuthenticator } from 'manekani-core';
import {
  Configuration,
  ConfigurationParameters,
  FrontendApi,
} from '@ory/client';

export type { ConfigurationParameters } from '@ory/client';

export class SsoAuthenticator implements ISsoAuthenticator {
  private ory: FrontendApi;

  constructor(options?: ConfigurationParameters) {
    this.ory = new FrontendApi(new Configuration(options));
  }

  requiredCookies(): string[] {
    return ['ory_kratos_session'];
  }

  async getCookieSession(cookies: string): Promise<UserSession> {
    try {
      const { data: orySession } = await this.ory.toSession({
        cookie: cookies,
      });

      return {
        userId: orySession.identity!.id,
        email: orySession.identity!.traits['email'],
      };
    } catch (error) {
      throw new UnauthorizedError({
        cause: error,
        description: 'Could not get the user session',
      });
    }
  }

  async registerUsername(_userId: string, _username: string): Promise<void> {
    throw new Error('Method not implemented.');
  }

  async updateUsername(_userId: string, _username: string): Promise<void> {
    throw new Error('Method not implemented.');
  }
}
