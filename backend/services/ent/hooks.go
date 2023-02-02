package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/hook"
	"sixels.io/manekani/ent/subject"
)

type IsSubject interface {
	ent.Mutation
	SetID(uuid.UUID)
	SetLevel(int32)
	Level() (int32, bool)
	ID() (uuid.UUID, bool)
}

func SetupHooks(client *ent.Client) {
	client.Radical.Use(hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.RadicalFunc(func(ctx context.Context, rm *ent.RadicalMutation) (ent.Value, error) {
				return createSubject(client, next, ctx, rm, subject.KindRadical)
			})
		},
		ent.OpCreate))
	client.Kanji.Use(hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.KanjiFunc(func(ctx context.Context, km *ent.KanjiMutation) (ent.Value, error) {
				return createSubject(client, next, ctx, km, subject.KindKanji)
			})
		},
		ent.OpCreate))
	client.Vocabulary.Use(hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.VocabularyFunc(func(ctx context.Context, vm *ent.VocabularyMutation) (ent.Value, error) {
				return createSubject(client, next, ctx, vm, subject.KindVocabulary)
			})
		},
		ent.OpCreate))

	client.Radical.Use(hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.RadicalFunc(func(ctx context.Context, rm *ent.RadicalMutation) (ent.Value, error) {
				return deleteSubject(client, next, ctx, rm)
			})
		},
		ent.OpDelete|ent.OpDeleteOne))
	client.Kanji.Use(hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.KanjiFunc(func(ctx context.Context, km *ent.KanjiMutation) (ent.Value, error) {
				return deleteSubject(client, next, ctx, km)
			})
		},
		ent.OpDelete|ent.OpDeleteOne))
	client.Vocabulary.Use(hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.VocabularyFunc(func(ctx context.Context, vm *ent.VocabularyMutation) (ent.Value, error) {
				return deleteSubject(client, next, ctx, vm)
			})
		},
		ent.OpDelete|ent.OpDeleteOne))
}

func createSubject(
	client *ent.Client,
	next ent.Mutator,
	ctx context.Context,
	mutation IsSubject,
	kind subject.Kind,
) (ent.Value, error) {
	id, isPresent := mutation.ID()
	level, _ := mutation.Level()

	if !isPresent {
		log.Println("Subject have no id defined, generating a new one")
		id = uuid.New()
		mutation.SetID(id)
	}

	if err := client.Subject.Create().
		SetID(id).
		SetLevel(level).
		SetKind(kind).
		Exec(ctx); err != nil {
		log.Printf("create subject error: %v\n", err)
		return nil, err
	}

	return next.Mutate(ctx, mutation)
}

func deleteSubject(
	client *ent.Client,
	next ent.Mutator,
	ctx context.Context,
	mutation IsSubject,
) (ent.Value, error) {
	id, isPresent := mutation.ID()

	if !isPresent {
		return nil, fmt.Errorf("subject have no id defined")
	}

	if err := client.Subject.DeleteOneID(id).
		Exec(ctx); err != nil {
		log.Printf("delete subject error: %v\n", err)
		return nil, err
	}

	return next.Mutate(ctx, mutation)
}
