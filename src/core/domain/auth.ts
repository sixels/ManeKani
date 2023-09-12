import { Static, Type } from '@sinclair/typebox';

export type UserSession = Static<typeof UserSessionSchema>;
export const UserSessionSchema = Type.Object({
  id: Type.String({ format: 'uuid' }),
  email: Type.Optional(Type.String({ format: 'email' })),
});
