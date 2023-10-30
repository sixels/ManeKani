import { Static, Type } from '@sinclair/typebox';

import { UuidSchema } from './common';

export type Deck = Static<typeof DeckSchema>;
export const DeckSchema = Type.Object({
  /**
   * The deck's unique identifier.
   */
  id: UuidSchema,
  /**
   * The deck's creation date.
   */
  createdAt: Type.Date(),
  /**
   * The deck's last update date.
   */
  updatedAt: Type.Date(),

  /**
   * The deck's name.
   * @example "WaniKani"
   */
  name: Type.String({ maxLength: 50 }),

  /**
   * The description about the deck.
   * @example "A deck for learning Japanese"
   */
  description: Type.String({ maxLength: 350 }),
  /**
   * The subjects that belong to this deck.
   */
  subjectIds: Type.Array(UuidSchema),
  /**
   * The deck's owner.
   */
  ownerId: Type.String(),
});

export type CreateDeckDto = Static<typeof CreateDeckSchema>;
export const CreateDeckSchema = Type.Pick(DeckSchema, ['name', 'description']);

export type UpdateDeckDto = Static<typeof UpdateDeckSchema>;
export const UpdateDeckSchema = Type.Pick(Type.Partial(DeckSchema), [
  'name',
  'description',
]);

export type GetDecksFilters = Static<typeof GetDecksFiltersSchema>;
export const GetDecksFiltersSchema = Type.Object({
  page: Type.Optional(Type.Integer({ minimum: 1 })),
  ids: Type.Optional(Type.Array(UuidSchema)),
  owners: Type.Optional(Type.Array(UuidSchema)),
  content: Type.Optional(Type.String()),
});
