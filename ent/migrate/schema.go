// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APITokensColumns holds the columns for the "api_tokens" table.
	APITokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeBytes},
		{Name: "name", Type: field.TypeString, Size: 20},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "frozen"}},
		{Name: "used_at", Type: field.TypeTime, Nullable: true},
		{Name: "token", Type: field.TypeString},
		{Name: "prefix", Type: field.TypeString},
		{Name: "claims", Type: field.TypeJSON},
		{Name: "user_api_tokens", Type: field.TypeString},
	}
	// APITokensTable holds the schema information for the "api_tokens" table.
	APITokensTable = &schema.Table{
		Name:       "api_tokens",
		Columns:    APITokensColumns,
		PrimaryKey: []*schema.Column{APITokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "api_tokens_users_api_tokens",
				Columns:    []*schema.Column{APITokensColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "apitoken_name_user_api_tokens",
				Unique:  true,
				Columns: []*schema.Column{APITokensColumns[1], APITokensColumns[7]},
			},
		},
	}
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "progress", Type: field.TypeUint8, Default: 0},
		{Name: "total_errors", Type: field.TypeInt32, Default: 0},
		{Name: "unlocked_at", Type: field.TypeTime, Nullable: true},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "passed_at", Type: field.TypeTime, Nullable: true},
		{Name: "available_at", Type: field.TypeTime, Nullable: true},
		{Name: "burned_at", Type: field.TypeTime, Nullable: true},
		{Name: "deck_progress_cards", Type: field.TypeInt},
		{Name: "subject_cards", Type: field.TypeUUID},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_deck_progresses_cards",
				Columns:    []*schema.Column{CardsColumns[10]},
				RefColumns: []*schema.Column{DeckProgressesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "cards_subjects_cards",
				Columns:    []*schema.Column{CardsColumns[11]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DecksColumns holds the columns for the "decks" table.
	DecksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "description", Type: field.TypeString, Size: 300},
		{Name: "user_decks", Type: field.TypeString},
	}
	// DecksTable holds the schema information for the "decks" table.
	DecksTable = &schema.Table{
		Name:       "decks",
		Columns:    DecksColumns,
		PrimaryKey: []*schema.Column{DecksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "decks_users_decks",
				Columns:    []*schema.Column{DecksColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DeckProgressesColumns holds the columns for the "deck_progresses" table.
	DeckProgressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "level", Type: field.TypeUint32, Default: 1},
		{Name: "deck_users_progress", Type: field.TypeUUID},
		{Name: "user_decks_progress", Type: field.TypeString},
	}
	// DeckProgressesTable holds the schema information for the "deck_progresses" table.
	DeckProgressesTable = &schema.Table{
		Name:       "deck_progresses",
		Columns:    DeckProgressesColumns,
		PrimaryKey: []*schema.Column{DeckProgressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deck_progresses_decks_users_progress",
				Columns:    []*schema.Column{DeckProgressesColumns[2]},
				RefColumns: []*schema.Column{DecksColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "deck_progresses_users_decks_progress",
				Columns:    []*schema.Column{DeckProgressesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "deckprogress_user_decks_progress_deck_users_progress",
				Unique:  true,
				Columns: []*schema.Column{DeckProgressesColumns[3], DeckProgressesColumns[2]},
			},
		},
	}
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "errors", Type: field.TypeJSON},
		{Name: "start_progress", Type: field.TypeUint8},
		{Name: "end_progress", Type: field.TypeUint8},
		{Name: "card_reviews", Type: field.TypeUUID},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reviews_cards_reviews",
				Columns:    []*schema.Column{ReviewsColumns[5]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "kind", Type: field.TypeString, Size: 2147483647},
		{Name: "level", Type: field.TypeInt32},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "value", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "value_image", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "slug", Type: field.TypeString, Size: 36},
		{Name: "priority", Type: field.TypeUint8},
		{Name: "resources", Type: field.TypeJSON, Nullable: true},
		{Name: "study_data", Type: field.TypeJSON},
		{Name: "additional_study_data", Type: field.TypeJSON, Nullable: true},
		{Name: "deck_subjects", Type: field.TypeUUID},
		{Name: "user_subjects", Type: field.TypeString},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subjects_decks_subjects",
				Columns:    []*schema.Column{SubjectsColumns[13]},
				RefColumns: []*schema.Column{DecksColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "subjects_users_subjects",
				Columns:    []*schema.Column{SubjectsColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "subject_kind_slug_deck_subjects",
				Unique:  true,
				Columns: []*schema.Column{SubjectsColumns[3], SubjectsColumns[8], SubjectsColumns[13]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "pending_actions", Type: field.TypeJSON, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 255},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// DeckSubscribersColumns holds the columns for the "deck_subscribers" table.
	DeckSubscribersColumns = []*schema.Column{
		{Name: "deck_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeString, Size: 36},
	}
	// DeckSubscribersTable holds the schema information for the "deck_subscribers" table.
	DeckSubscribersTable = &schema.Table{
		Name:       "deck_subscribers",
		Columns:    DeckSubscribersColumns,
		PrimaryKey: []*schema.Column{DeckSubscribersColumns[0], DeckSubscribersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deck_subscribers_deck_id",
				Columns:    []*schema.Column{DeckSubscribersColumns[0]},
				RefColumns: []*schema.Column{DecksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "deck_subscribers_user_id",
				Columns:    []*schema.Column{DeckSubscribersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubjectSimilarColumns holds the columns for the "subject_similar" table.
	SubjectSimilarColumns = []*schema.Column{
		{Name: "subject_id", Type: field.TypeUUID},
		{Name: "similar_id", Type: field.TypeUUID},
	}
	// SubjectSimilarTable holds the schema information for the "subject_similar" table.
	SubjectSimilarTable = &schema.Table{
		Name:       "subject_similar",
		Columns:    SubjectSimilarColumns,
		PrimaryKey: []*schema.Column{SubjectSimilarColumns[0], SubjectSimilarColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subject_similar_subject_id",
				Columns:    []*schema.Column{SubjectSimilarColumns[0]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "subject_similar_similar_id",
				Columns:    []*schema.Column{SubjectSimilarColumns[1]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubjectDependentsColumns holds the columns for the "subject_dependents" table.
	SubjectDependentsColumns = []*schema.Column{
		{Name: "subject_id", Type: field.TypeUUID},
		{Name: "dependency_id", Type: field.TypeUUID},
	}
	// SubjectDependentsTable holds the schema information for the "subject_dependents" table.
	SubjectDependentsTable = &schema.Table{
		Name:       "subject_dependents",
		Columns:    SubjectDependentsColumns,
		PrimaryKey: []*schema.Column{SubjectDependentsColumns[0], SubjectDependentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subject_dependents_subject_id",
				Columns:    []*schema.Column{SubjectDependentsColumns[0]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "subject_dependents_dependency_id",
				Columns:    []*schema.Column{SubjectDependentsColumns[1]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APITokensTable,
		CardsTable,
		DecksTable,
		DeckProgressesTable,
		ReviewsTable,
		SubjectsTable,
		UsersTable,
		DeckSubscribersTable,
		SubjectSimilarTable,
		SubjectDependentsTable,
	}
)

func init() {
	APITokensTable.ForeignKeys[0].RefTable = UsersTable
	CardsTable.ForeignKeys[0].RefTable = DeckProgressesTable
	CardsTable.ForeignKeys[1].RefTable = SubjectsTable
	DecksTable.ForeignKeys[0].RefTable = UsersTable
	DeckProgressesTable.ForeignKeys[0].RefTable = DecksTable
	DeckProgressesTable.ForeignKeys[1].RefTable = UsersTable
	ReviewsTable.ForeignKeys[0].RefTable = CardsTable
	SubjectsTable.ForeignKeys[0].RefTable = DecksTable
	SubjectsTable.ForeignKeys[1].RefTable = UsersTable
	DeckSubscribersTable.ForeignKeys[0].RefTable = DecksTable
	DeckSubscribersTable.ForeignKeys[1].RefTable = UsersTable
	SubjectSimilarTable.ForeignKeys[0].RefTable = SubjectsTable
	SubjectSimilarTable.ForeignKeys[1].RefTable = SubjectsTable
	SubjectDependentsTable.ForeignKeys[0].RefTable = SubjectsTable
	SubjectDependentsTable.ForeignKeys[1].RefTable = SubjectsTable
}
