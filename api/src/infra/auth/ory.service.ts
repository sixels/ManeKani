import { UserSession } from 'manekani-core';
import { UnauthorizedError } from 'manekani-core';
import { ISsoAuthenticator } from 'manekani-core';
import { Inject, Injectable } from '@nestjs/common';
import {
  Configuration,
  ConfigurationParameters,
  FrontendApi,
} from '@ory/client';

export const ORY_OPTIONS_KEY = 'ORY_CLIENT_OPTIONS';

@Injectable()
export class OryService implements ISsoAuthenticator {
  private ory: FrontendApi;

  constructor(@Inject(ORY_OPTIONS_KEY) options: ConfigurationParameters) {
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
        userId: orySession.identity.id,
        email: orySession.identity.traits['email'],
      };
    } catch (error) {
      throw new UnauthorizedError({
        cause: error,
        description: 'Could not get the user session',
      });
    }
  }
}
