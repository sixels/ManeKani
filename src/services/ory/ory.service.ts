import { UserSession } from '@/core/domain/auth';
import { UnauthorizedError } from '@/core/domain/error';
import { ISsoAuthenticator } from '@/core/ports/auth';
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

  async getSessionFromCookies(cookies: string): Promise<UserSession> {
    try {
      const { data: orySession } = await this.ory.toSession({
        cookie: cookies,
      });

      return {
        id: orySession.identity.id,
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
