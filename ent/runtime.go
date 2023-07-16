// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	ulid "github.com/oklog/ulid/v2"
	"github.com/sixels/manekani/ent/apitoken"
	"github.com/sixels/manekani/ent/card"
	"github.com/sixels/manekani/ent/deck"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/review"
	"github.com/sixels/manekani/ent/schema"
	"github.com/sixels/manekani/ent/subject"
	"github.com/sixels/manekani/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	apitokenFields := schema.ApiToken{}.Fields()
	_ = apitokenFields
	// apitokenDescName is the schema descriptor for name field.
	apitokenDescName := apitokenFields[1].Descriptor()
	// apitoken.NameValidator is a validator for the "name" field. It is called by the builders before save.
	apitoken.NameValidator = func() func(string) error {
		validators := apitokenDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// apitokenDescID is the schema descriptor for id field.
	apitokenDescID := apitokenFields[0].Descriptor()
	// apitoken.DefaultID holds the default value on creation for the id field.
	apitoken.DefaultID = apitokenDescID.Default.(func() ulid.ULID)
	cardMixin := schema.Card{}.Mixin()
	cardMixinFields0 := cardMixin[0].Fields()
	_ = cardMixinFields0
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescCreatedAt is the schema descriptor for created_at field.
	cardDescCreatedAt := cardMixinFields0[0].Descriptor()
	// card.DefaultCreatedAt holds the default value on creation for the created_at field.
	card.DefaultCreatedAt = cardDescCreatedAt.Default.(func() time.Time)
	// cardDescUpdatedAt is the schema descriptor for updated_at field.
	cardDescUpdatedAt := cardMixinFields0[1].Descriptor()
	// card.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	card.DefaultUpdatedAt = cardDescUpdatedAt.Default.(func() time.Time)
	// card.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	card.UpdateDefaultUpdatedAt = cardDescUpdatedAt.UpdateDefault.(func() time.Time)
	// cardDescProgress is the schema descriptor for progress field.
	cardDescProgress := cardFields[1].Descriptor()
	// card.DefaultProgress holds the default value on creation for the progress field.
	card.DefaultProgress = cardDescProgress.Default.(uint8)
	// cardDescTotalErrors is the schema descriptor for total_errors field.
	cardDescTotalErrors := cardFields[2].Descriptor()
	// card.DefaultTotalErrors holds the default value on creation for the total_errors field.
	card.DefaultTotalErrors = cardDescTotalErrors.Default.(int32)
	// cardDescID is the schema descriptor for id field.
	cardDescID := cardFields[0].Descriptor()
	// card.DefaultID holds the default value on creation for the id field.
	card.DefaultID = cardDescID.Default.(func() uuid.UUID)
	deckMixin := schema.Deck{}.Mixin()
	deckMixinFields0 := deckMixin[0].Fields()
	_ = deckMixinFields0
	deckFields := schema.Deck{}.Fields()
	_ = deckFields
	// deckDescCreatedAt is the schema descriptor for created_at field.
	deckDescCreatedAt := deckMixinFields0[0].Descriptor()
	// deck.DefaultCreatedAt holds the default value on creation for the created_at field.
	deck.DefaultCreatedAt = deckDescCreatedAt.Default.(func() time.Time)
	// deckDescUpdatedAt is the schema descriptor for updated_at field.
	deckDescUpdatedAt := deckMixinFields0[1].Descriptor()
	// deck.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deck.DefaultUpdatedAt = deckDescUpdatedAt.Default.(func() time.Time)
	// deck.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deck.UpdateDefaultUpdatedAt = deckDescUpdatedAt.UpdateDefault.(func() time.Time)
	// deckDescName is the schema descriptor for name field.
	deckDescName := deckFields[1].Descriptor()
	// deck.NameValidator is a validator for the "name" field. It is called by the builders before save.
	deck.NameValidator = deckDescName.Validators[0].(func(string) error)
	// deckDescDescription is the schema descriptor for description field.
	deckDescDescription := deckFields[2].Descriptor()
	// deck.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	deck.DescriptionValidator = deckDescDescription.Validators[0].(func(string) error)
	// deckDescID is the schema descriptor for id field.
	deckDescID := deckFields[0].Descriptor()
	// deck.DefaultID holds the default value on creation for the id field.
	deck.DefaultID = deckDescID.Default.(func() uuid.UUID)
	deckprogressFields := schema.DeckProgress{}.Fields()
	_ = deckprogressFields
	// deckprogressDescLevel is the schema descriptor for level field.
	deckprogressDescLevel := deckprogressFields[0].Descriptor()
	// deckprogress.DefaultLevel holds the default value on creation for the level field.
	deckprogress.DefaultLevel = deckprogressDescLevel.Default.(uint32)
	// deckprogress.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	deckprogress.LevelValidator = deckprogressDescLevel.Validators[0].(func(uint32) error)
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescCreatedAt is the schema descriptor for created_at field.
	reviewDescCreatedAt := reviewFields[1].Descriptor()
	// review.DefaultCreatedAt holds the default value on creation for the created_at field.
	review.DefaultCreatedAt = reviewDescCreatedAt.Default.(func() time.Time)
	// reviewDescID is the schema descriptor for id field.
	reviewDescID := reviewFields[0].Descriptor()
	// review.DefaultID holds the default value on creation for the id field.
	review.DefaultID = reviewDescID.Default.(func() uuid.UUID)
	subjectMixin := schema.Subject{}.Mixin()
	subjectMixinFields0 := subjectMixin[0].Fields()
	_ = subjectMixinFields0
	subjectFields := schema.Subject{}.Fields()
	_ = subjectFields
	// subjectDescCreatedAt is the schema descriptor for created_at field.
	subjectDescCreatedAt := subjectMixinFields0[0].Descriptor()
	// subject.DefaultCreatedAt holds the default value on creation for the created_at field.
	subject.DefaultCreatedAt = subjectDescCreatedAt.Default.(func() time.Time)
	// subjectDescUpdatedAt is the schema descriptor for updated_at field.
	subjectDescUpdatedAt := subjectMixinFields0[1].Descriptor()
	// subject.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	subject.DefaultUpdatedAt = subjectDescUpdatedAt.Default.(func() time.Time)
	// subject.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	subject.UpdateDefaultUpdatedAt = subjectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// subjectDescLevel is the schema descriptor for level field.
	subjectDescLevel := subjectFields[2].Descriptor()
	// subject.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	subject.LevelValidator = subjectDescLevel.Validators[0].(func(int32) error)
	// subjectDescName is the schema descriptor for name field.
	subjectDescName := subjectFields[3].Descriptor()
	// subject.NameValidator is a validator for the "name" field. It is called by the builders before save.
	subject.NameValidator = subjectDescName.Validators[0].(func(string) error)
	// subjectDescValue is the schema descriptor for value field.
	subjectDescValue := subjectFields[4].Descriptor()
	// subject.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	subject.ValueValidator = subjectDescValue.Validators[0].(func(string) error)
	// subjectDescSlug is the schema descriptor for slug field.
	subjectDescSlug := subjectFields[6].Descriptor()
	// subject.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	subject.SlugValidator = func() func(string) error {
		validators := subjectDescSlug.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(slug string) error {
			for _, fn := range fns {
				if err := fn(slug); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// subjectDescID is the schema descriptor for id field.
	subjectDescID := subjectFields[0].Descriptor()
	// subject.DefaultID holds the default value on creation for the id field.
	subject.DefaultID = subjectDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() string)
}
