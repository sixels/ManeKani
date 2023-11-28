import { Static, Type } from "@sinclair/typebox";
import { TypeSlug, UuidSchema } from "./common";

const PrimitiveSchema = Type.Union([
	Type.String({ maxLength: 80 }),
	Type.Number(),
	Type.Boolean(),
]);

export type StudyDataItem = Static<typeof StudyDataItemSchema>;
export const StudyDataItemSchema = Type.Object({
	/**
	 * The value of this study data item.
	 * @example "one"
	 */
	value: Type.String({ maxLength: 50 }),
	/**
	 * Whether this item is the primary answer or not.
	 */
	isPrimary: Type.Boolean(),
	/**
	 * Whether this item is a valid answer or not.
	 */
	isValidAnswer: Type.Boolean(),
	/**
	 * Whether this item should be hidden or not. This may still be a valid
	 * answer, but not an expected answer.
	 */
	isHidden: Type.Boolean(),
	/**
	 * The category of this item.
	 * @example "onyomi"
	 * @example "kunyomi"
	 * @example "nanori"
	 */
	category: Type.Optional(Type.String({ maxLength: 30 })),
	/**
	 * The id of the resource related to this answer.
	 * @example "resource/KyzjB4BoUr"
	 */
	resourceId: Type.Optional(Type.String({ maxLength: 30 })),
});

export type StudyData = Static<typeof StudyDataSchema>;
export const StudyDataSchema = Type.Object({
	/**
	 * The type of this study data.
	 * @example "meaning"
	 * @example "reading"
	 */
	type: Type.String({ maxLength: 30 }),
	/**
	 * A mnemonic to help the user remember the subject's value.
	 * @example "Lying on the <radical>ground</radical> is something that looks \
	 * just like the ground, the number One. Why is this One lying down? It's \
	 * been shot by the number two. It's lying there, bleeding out and dying. The \
	 * number <kanji>One</kanji> doesn't have long to live."
	 */
	mnemonic: Type.Optional(Type.String({ maxLength: 300 })),
	/**
	 * The study data items which will be used to validate the user's answer.
	 * @example [{
	 * value: "one",
	 * isPrimary: true,
	 * isValidAnswer: true,
	 * isHidden: false,
	 * }]
	 * @example [{
	 * value: "いち",
	 * isPrimary: true,
	 * isValidAnswer: true,
	 * isHidden: false,
	 * category: "onyomi",
	 * }, {
	 * value: "いっ",
	 * isPrimary: false,
	 * isValidAnswer: true,
	 * isHidden: false,
	 * category: "onyomi",
	 * }]
	 */
	items: Type.Array(StudyDataItemSchema),
});

export type Resource = Static<typeof ResourceSchema>;
export const ResourceSchema = Type.Object({
	/**
	 * The resource name
	 * @example "pronunciation audio"
	 */
	name: Type.String({ maxLength: 30 }),
	/**
	 * The resource data
	 * @example "/<random-hash>.ogg"
	 */
	data: Type.Optional(Type.String({ maxLength: 200 })),
	/**
	 * Additional metadata to the resource
	 * @example {"type": "audio/ogg", "voiceActor": "John Doe"}
	 * @example {
	 *  "sentences": [{"en": "one", "pt": "um"}, {"en": "two", "pt": "dois"}],
	 *  "patterns": {"おやつに~": [{ "en": "...", "pt": "..."}]},
	 * }
	 */
	metadata: Type.Record(
		Type.String(),
		Type.Union([
			PrimitiveSchema,

			Type.Array(PrimitiveSchema),
			Type.Array(Type.Record(Type.String(), PrimitiveSchema)),

			Type.Record(Type.String(), PrimitiveSchema),
			Type.Record(Type.String(), Type.Array(PrimitiveSchema)),
			Type.Record(
				Type.String(),
				Type.Array(Type.Record(Type.String(), PrimitiveSchema)),
			),
		]),
		{ default: {} },
	),
});

export type Subject = Static<typeof SubjectSchema>;
export const SubjectSchema = Type.Object({
	/**
	 * The subject unique identifier.
	 */
	id: UuidSchema,
	/**
	 * The date when the subject was created.
	 */
	createdAt: Type.Date(),
	/**
	 * The date when the subject was last updated.
	 */
	updatedAt: Type.Date(),
	/**
	 * The card category.
	 * @example "kanji"
	 */
	category: Type.String({ maxLength: 30, minLength: 1 }),
	/**
	 * The level of difficult of the subject. should be greater or equal than 1.
	 */
	level: Type.Integer({ minimum: 1, maximum: 1000 }),
	/**
	 * The subject's name.
	 */
	name: Type.String({ maxLength: 50, minLength: 1 }),
	/**
	 * The subject's value which will be displayed at the front of the card.
	 * @example "一" (which would be the value of the kanji subject "one")
	 */
	value: Type.Optional(Type.String({ maxLength: 150, minLength: 1 })),
	/**
	 * In addition to the text value, the subject can also have an image to be
	 * displayed at the front of the card. This value shoul be a valid image id
	 * generated when a user uploads an image through the files api.
	 * @example "subject/KyzjB4BoUr"
	 */
	valueImage: Type.Optional(Type.String({ maxLength: 40 })),
	/**
	 * The subject's slug which will be used to generate the subject's url
	 * @example "one"
	 */
	slug: TypeSlug({
		maxLength: 50,
		minLength: 1,
	}),
	/**
	 * The subject's priority which will be used as precedence for this subject to
	 * show up at the study session. The lower the number, higher the priority (the
	 * minimum number is 0).
	 * @example 1
	 */
	priority: Type.Integer({ minimum: 0, maximum: 255 }),
	/**
	 * The subject's resources which will be used to validate the user's answer.
	 * The mnemonic field should be a valid markdown string.
	 * @example [{
	 *  type: "meaning",
	 *  mnemonic: "Lying on the <radical>ground</radical> is something that looks \
	 *  just like the ground, the number One. Why is this One lying down? It's \
	 *  been shot by the number two. It's lying there, bleeding out and dying. The \
	 *  number <kanji>One</kanji> doesn't have long to live.",
	 *  items: [
	 *    {value: "one", isPrimary: true, isValidAnswer: true, isHidden: false},
	 *    ...
	 *  ]
	 *  }, ...]
	 */
	studyData: Type.Array(StudyDataSchema, { default: [] }),
	/**
	 * The same as `studyData` but only for text resources. can be used for this
	 * like links, pronunciation, patterns and more.
	 */
	resources: Type.Array(ResourceSchema, { default: [] }),
	/**
	 * The subject dependencies that needs to be completed before unlocking this subject.
	 */
	dependencies: Type.Array(UuidSchema, { default: [] }),
	/**
	 * The subject dependents that might be unlocked after this subject is completed.
	 */
	dependents: Type.Array(UuidSchema, { default: [] }),
	/**
	 * Similar subjects to help the user not confuse this subject with another one.
	 */
	similar: Type.Array(UuidSchema, { default: [] }),
	/**
	 * The deck in which this subject belongs to.
	 */
	deckId: UuidSchema,
	/**
	 * The user that owns this subject.
	 */
	ownerId: Type.String(),
});

export type PartialSubject = Static<typeof PartialSubjectSchema>;
export const PartialSubjectSchema = Type.Pick(SubjectSchema, [
	"id",
	"createdAt",
	"updatedAt",
	"category",
	"level",
	"name",
	"value",
	"valueImage",
	"slug",
	"priority",
	"dependencies",
	"dependents",
	"similar",
	"deckId",
	"ownerId",
]);

export type CreateSubjectDto = Static<typeof CreateSubjectSchema>;
export const CreateSubjectSchema = Type.Pick(SubjectSchema, [
	"category",
	"level",
	"name",
	"value",
	"valueImage",
	"slug",
	"priority",
	"studyData",
	"resources",
	"dependencies",
	"dependents",
	"similar",
]);

export type UpdateSubjectDto = Static<typeof UpdateSubjectSchema>;
export const UpdateSubjectSchema = Type.Partial(CreateSubjectSchema);

export type GetSubjectsFilters = Static<typeof GetSubjectsFiltersSchema>;
export const GetSubjectsFiltersSchema = Type.Object({
	page: Type.Optional(Type.Integer({ minimum: 1 })),
	ids: Type.Optional(Type.Array(UuidSchema)),
	categories: Type.Optional(Type.Array(Type.String())),
	levels: Type.Optional(Type.Array(Type.Integer({ minimum: 1 }))),
	slugs: Type.Optional(Type.Array(Type.String())),
	decks: Type.Optional(Type.Array(UuidSchema)),
	owners: Type.Optional(Type.Array(UuidSchema)),
});
