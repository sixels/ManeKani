// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"sixels.io/manekani/ent/kanji"
	"sixels.io/manekani/ent/radical"
	"sixels.io/manekani/ent/schema"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/ent/vocabulary"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	kanjiMixin := schema.Kanji{}.Mixin()
	kanjiMixinFields0 := kanjiMixin[0].Fields()
	_ = kanjiMixinFields0
	kanjiFields := schema.Kanji{}.Fields()
	_ = kanjiFields
	// kanjiDescCreatedAt is the schema descriptor for created_at field.
	kanjiDescCreatedAt := kanjiMixinFields0[0].Descriptor()
	// kanji.DefaultCreatedAt holds the default value on creation for the created_at field.
	kanji.DefaultCreatedAt = kanjiDescCreatedAt.Default.(func() time.Time)
	// kanjiDescUpdatedAt is the schema descriptor for updated_at field.
	kanjiDescUpdatedAt := kanjiMixinFields0[1].Descriptor()
	// kanji.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	kanji.DefaultUpdatedAt = kanjiDescUpdatedAt.Default.(func() time.Time)
	// kanji.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	kanji.UpdateDefaultUpdatedAt = kanjiDescUpdatedAt.UpdateDefault.(func() time.Time)
	// kanjiDescSymbol is the schema descriptor for symbol field.
	kanjiDescSymbol := kanjiFields[1].Descriptor()
	// kanji.SymbolValidator is a validator for the "symbol" field. It is called by the builders before save.
	kanji.SymbolValidator = func() func(string) error {
		validators := kanjiDescSymbol.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(symbol string) error {
			for _, fn := range fns {
				if err := fn(symbol); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// kanjiDescName is the schema descriptor for name field.
	kanjiDescName := kanjiFields[2].Descriptor()
	// kanji.NameValidator is a validator for the "name" field. It is called by the builders before save.
	kanji.NameValidator = kanjiDescName.Validators[0].(func(string) error)
	// kanjiDescLevel is the schema descriptor for level field.
	kanjiDescLevel := kanjiFields[5].Descriptor()
	// kanji.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	kanji.LevelValidator = kanjiDescLevel.Validators[0].(func(int32) error)
	// kanjiDescReading is the schema descriptor for reading field.
	kanjiDescReading := kanjiFields[6].Descriptor()
	// kanji.ReadingValidator is a validator for the "reading" field. It is called by the builders before save.
	kanji.ReadingValidator = kanjiDescReading.Validators[0].(func(string) error)
	// kanjiDescMeaningMnemonic is the schema descriptor for meaning_mnemonic field.
	kanjiDescMeaningMnemonic := kanjiFields[10].Descriptor()
	// kanji.MeaningMnemonicValidator is a validator for the "meaning_mnemonic" field. It is called by the builders before save.
	kanji.MeaningMnemonicValidator = kanjiDescMeaningMnemonic.Validators[0].(func(string) error)
	// kanjiDescReadingMnemonic is the schema descriptor for reading_mnemonic field.
	kanjiDescReadingMnemonic := kanjiFields[11].Descriptor()
	// kanji.ReadingMnemonicValidator is a validator for the "reading_mnemonic" field. It is called by the builders before save.
	kanji.ReadingMnemonicValidator = kanjiDescReadingMnemonic.Validators[0].(func(string) error)
	// kanjiDescID is the schema descriptor for id field.
	kanjiDescID := kanjiFields[0].Descriptor()
	// kanji.DefaultID holds the default value on creation for the id field.
	kanji.DefaultID = kanjiDescID.Default.(func() uuid.UUID)
	radicalMixin := schema.Radical{}.Mixin()
	radicalMixinFields0 := radicalMixin[0].Fields()
	_ = radicalMixinFields0
	radicalFields := schema.Radical{}.Fields()
	_ = radicalFields
	// radicalDescCreatedAt is the schema descriptor for created_at field.
	radicalDescCreatedAt := radicalMixinFields0[0].Descriptor()
	// radical.DefaultCreatedAt holds the default value on creation for the created_at field.
	radical.DefaultCreatedAt = radicalDescCreatedAt.Default.(func() time.Time)
	// radicalDescUpdatedAt is the schema descriptor for updated_at field.
	radicalDescUpdatedAt := radicalMixinFields0[1].Descriptor()
	// radical.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	radical.DefaultUpdatedAt = radicalDescUpdatedAt.Default.(func() time.Time)
	// radical.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	radical.UpdateDefaultUpdatedAt = radicalDescUpdatedAt.UpdateDefault.(func() time.Time)
	// radicalDescName is the schema descriptor for name field.
	radicalDescName := radicalFields[1].Descriptor()
	// radical.NameValidator is a validator for the "name" field. It is called by the builders before save.
	radical.NameValidator = radicalDescName.Validators[0].(func(string) error)
	// radicalDescLevel is the schema descriptor for level field.
	radicalDescLevel := radicalFields[2].Descriptor()
	// radical.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	radical.LevelValidator = radicalDescLevel.Validators[0].(func(int32) error)
	// radicalDescSymbol is the schema descriptor for symbol field.
	radicalDescSymbol := radicalFields[3].Descriptor()
	// radical.SymbolValidator is a validator for the "symbol" field. It is called by the builders before save.
	radical.SymbolValidator = radicalDescSymbol.Validators[0].(func(string) error)
	// radicalDescMeaningMnemonic is the schema descriptor for meaning_mnemonic field.
	radicalDescMeaningMnemonic := radicalFields[4].Descriptor()
	// radical.MeaningMnemonicValidator is a validator for the "meaning_mnemonic" field. It is called by the builders before save.
	radical.MeaningMnemonicValidator = radicalDescMeaningMnemonic.Validators[0].(func(string) error)
	// radicalDescID is the schema descriptor for id field.
	radicalDescID := radicalFields[0].Descriptor()
	// radical.DefaultID holds the default value on creation for the id field.
	radical.DefaultID = radicalDescID.Default.(func() uuid.UUID)
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
	userDescEmail := userFields[2].Descriptor()
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
	vocabularyMixin := schema.Vocabulary{}.Mixin()
	vocabularyMixinFields0 := vocabularyMixin[0].Fields()
	_ = vocabularyMixinFields0
	vocabularyFields := schema.Vocabulary{}.Fields()
	_ = vocabularyFields
	// vocabularyDescCreatedAt is the schema descriptor for created_at field.
	vocabularyDescCreatedAt := vocabularyMixinFields0[0].Descriptor()
	// vocabulary.DefaultCreatedAt holds the default value on creation for the created_at field.
	vocabulary.DefaultCreatedAt = vocabularyDescCreatedAt.Default.(func() time.Time)
	// vocabularyDescUpdatedAt is the schema descriptor for updated_at field.
	vocabularyDescUpdatedAt := vocabularyMixinFields0[1].Descriptor()
	// vocabulary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vocabulary.DefaultUpdatedAt = vocabularyDescUpdatedAt.Default.(func() time.Time)
	// vocabulary.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vocabulary.UpdateDefaultUpdatedAt = vocabularyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// vocabularyDescName is the schema descriptor for name field.
	vocabularyDescName := vocabularyFields[1].Descriptor()
	// vocabulary.NameValidator is a validator for the "name" field. It is called by the builders before save.
	vocabulary.NameValidator = vocabularyDescName.Validators[0].(func(string) error)
	// vocabularyDescLevel is the schema descriptor for level field.
	vocabularyDescLevel := vocabularyFields[3].Descriptor()
	// vocabulary.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	vocabulary.LevelValidator = vocabularyDescLevel.Validators[0].(func(int32) error)
	// vocabularyDescWord is the schema descriptor for word field.
	vocabularyDescWord := vocabularyFields[4].Descriptor()
	// vocabulary.WordValidator is a validator for the "word" field. It is called by the builders before save.
	vocabulary.WordValidator = vocabularyDescWord.Validators[0].(func(string) error)
	// vocabularyDescReading is the schema descriptor for reading field.
	vocabularyDescReading := vocabularyFields[6].Descriptor()
	// vocabulary.ReadingValidator is a validator for the "reading" field. It is called by the builders before save.
	vocabulary.ReadingValidator = vocabularyDescReading.Validators[0].(func(string) error)
	// vocabularyDescMeaningMnemonic is the schema descriptor for meaning_mnemonic field.
	vocabularyDescMeaningMnemonic := vocabularyFields[10].Descriptor()
	// vocabulary.MeaningMnemonicValidator is a validator for the "meaning_mnemonic" field. It is called by the builders before save.
	vocabulary.MeaningMnemonicValidator = vocabularyDescMeaningMnemonic.Validators[0].(func(string) error)
	// vocabularyDescReadingMnemonic is the schema descriptor for reading_mnemonic field.
	vocabularyDescReadingMnemonic := vocabularyFields[11].Descriptor()
	// vocabulary.ReadingMnemonicValidator is a validator for the "reading_mnemonic" field. It is called by the builders before save.
	vocabulary.ReadingMnemonicValidator = vocabularyDescReadingMnemonic.Validators[0].(func(string) error)
	// vocabularyDescID is the schema descriptor for id field.
	vocabularyDescID := vocabularyFields[0].Descriptor()
	// vocabulary.DefaultID holds the default value on creation for the id field.
	vocabulary.DefaultID = vocabularyDescID.Default.(func() uuid.UUID)
}
