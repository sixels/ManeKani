import { Static, Type } from '@sinclair/typebox';

export const UsernameSchema = Type.String({
  maxLength: 25,
  pattern: '^(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$',
});

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
  username: Type.Optional(UsernameSchema),
  /**
   * The user's display name
   */
  displayName: Type.Optional(Type.String({ maxLength: 50 })),

  /**
   * If this is a verified user
   */
  isVerified: Type.Boolean(),
  /**
   * If this user has completed their profile
   */
  isComplete: Type.Boolean(),

  /**
   * When the user account was created
   */
  createdAt: Type.String({ format: 'date-time' }),
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
export const CreateUserSchema = Type.Intersect([
  Type.Object({
    username: Type.Optional(UsernameSchema),
  }),
  Type.Pick(UserSchema, ['id', 'email', 'createdAt']),
]);

export type UpdateUserDto = Static<typeof UpdateUserSchema>;
export const UpdateUserSchema = Type.Partial(
  Type.Pick(UserSchema, ['username', 'displayName', 'isComplete']),
);
