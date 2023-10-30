import { Static, Type } from '@sinclair/typebox';

export type User = Static<typeof UserSchema>;
export const UserSchema = Type.Object({
  /**
   * The user's id
   */
  id: Type.String(),

  /**
   * The user's email
   */
  email: Type.String({ format: 'email' }),
  /**
   * The user's unique username
   */
  username: Type.String({
    maxLength: 25,
    pattern: '^(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$',
  }),

  /**
   * If this is a verified user
   */
  isVerified: Type.Boolean(),

  /**
   * When the user account was created
   */
  createdAt: Type.Date(),
  /**
   * When the user account was last updated
   */
  updatedAt: Type.Date(),
});

export type PublicUser = Static<typeof PublicUserSchema>;
export const PublicUserSchema = Type.Pick(UserSchema, [
  'username',
  'isVerified',
]);

export type CreateUserDto = Static<typeof CreateUserSchema>;
export const CreateUserSchema = Type.Pick(UserSchema, [
  'id',
  'username',
  'email',
  'createdAt',
]);
