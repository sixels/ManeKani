import { Static, Type } from '@sinclair/typebox';

import { UuidSchema } from './common';

export const DeckSchema = Type.Object({
  /**
   * The deck's unique identifier.
   */
  id: UuidSchema,
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
  subjects: Type.Array(UuidSchema),
  /**
   * The deck's owner.
   */
  ownerId: Type.String(),
});

export type Deck = Static<typeof DeckSchema>;
