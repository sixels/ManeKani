// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"sixels.io/manekani/ent/kanji"
	"sixels.io/manekani/ent/radical"
	"sixels.io/manekani/ent/vocabulary"
)

// KanjiCreate is the builder for creating a Kanji entity.
type KanjiCreate struct {
	config
	mutation *KanjiMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (kc *KanjiCreate) SetCreatedAt(t time.Time) *KanjiCreate {
	kc.mutation.SetCreatedAt(t)
	return kc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kc *KanjiCreate) SetNillableCreatedAt(t *time.Time) *KanjiCreate {
	if t != nil {
		kc.SetCreatedAt(*t)
	}
	return kc
}

// SetUpdatedAt sets the "updated_at" field.
func (kc *KanjiCreate) SetUpdatedAt(t time.Time) *KanjiCreate {
	kc.mutation.SetUpdatedAt(t)
	return kc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kc *KanjiCreate) SetNillableUpdatedAt(t *time.Time) *KanjiCreate {
	if t != nil {
		kc.SetUpdatedAt(*t)
	}
	return kc
}

// SetSymbol sets the "symbol" field.
func (kc *KanjiCreate) SetSymbol(s string) *KanjiCreate {
	kc.mutation.SetSymbol(s)
	return kc
}

// SetName sets the "name" field.
func (kc *KanjiCreate) SetName(s string) *KanjiCreate {
	kc.mutation.SetName(s)
	return kc
}

// SetAltNames sets the "alt_names" field.
func (kc *KanjiCreate) SetAltNames(pa pgtype.TextArray) *KanjiCreate {
	kc.mutation.SetAltNames(pa)
	return kc
}

// SetNillableAltNames sets the "alt_names" field if the given value is not nil.
func (kc *KanjiCreate) SetNillableAltNames(pa *pgtype.TextArray) *KanjiCreate {
	if pa != nil {
		kc.SetAltNames(*pa)
	}
	return kc
}

// SetSimilar sets the "similar" field.
func (kc *KanjiCreate) SetSimilar(pa pgtype.TextArray) *KanjiCreate {
	kc.mutation.SetSimilar(pa)
	return kc
}

// SetNillableSimilar sets the "similar" field if the given value is not nil.
func (kc *KanjiCreate) SetNillableSimilar(pa *pgtype.TextArray) *KanjiCreate {
	if pa != nil {
		kc.SetSimilar(*pa)
	}
	return kc
}

// SetLevel sets the "level" field.
func (kc *KanjiCreate) SetLevel(i int32) *KanjiCreate {
	kc.mutation.SetLevel(i)
	return kc
}

// SetReading sets the "reading" field.
func (kc *KanjiCreate) SetReading(s string) *KanjiCreate {
	kc.mutation.SetReading(s)
	return kc
}

// SetOnyomi sets the "onyomi" field.
func (kc *KanjiCreate) SetOnyomi(pa pgtype.TextArray) *KanjiCreate {
	kc.mutation.SetOnyomi(pa)
	return kc
}

// SetKunyomi sets the "kunyomi" field.
func (kc *KanjiCreate) SetKunyomi(pa pgtype.TextArray) *KanjiCreate {
	kc.mutation.SetKunyomi(pa)
	return kc
}

// SetNanori sets the "nanori" field.
func (kc *KanjiCreate) SetNanori(pa pgtype.TextArray) *KanjiCreate {
	kc.mutation.SetNanori(pa)
	return kc
}

// SetMeaningMnemonic sets the "meaning_mnemonic" field.
func (kc *KanjiCreate) SetMeaningMnemonic(s string) *KanjiCreate {
	kc.mutation.SetMeaningMnemonic(s)
	return kc
}

// SetReadingMnemonic sets the "reading_mnemonic" field.
func (kc *KanjiCreate) SetReadingMnemonic(s string) *KanjiCreate {
	kc.mutation.SetReadingMnemonic(s)
	return kc
}

// SetID sets the "id" field.
func (kc *KanjiCreate) SetID(u uuid.UUID) *KanjiCreate {
	kc.mutation.SetID(u)
	return kc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (kc *KanjiCreate) SetNillableID(u *uuid.UUID) *KanjiCreate {
	if u != nil {
		kc.SetID(*u)
	}
	return kc
}

// AddVocabularyIDs adds the "vocabularies" edge to the Vocabulary entity by IDs.
func (kc *KanjiCreate) AddVocabularyIDs(ids ...uuid.UUID) *KanjiCreate {
	kc.mutation.AddVocabularyIDs(ids...)
	return kc
}

// AddVocabularies adds the "vocabularies" edges to the Vocabulary entity.
func (kc *KanjiCreate) AddVocabularies(v ...*Vocabulary) *KanjiCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return kc.AddVocabularyIDs(ids...)
}

// AddRadicalIDs adds the "radicals" edge to the Radical entity by IDs.
func (kc *KanjiCreate) AddRadicalIDs(ids ...uuid.UUID) *KanjiCreate {
	kc.mutation.AddRadicalIDs(ids...)
	return kc
}

// AddRadicals adds the "radicals" edges to the Radical entity.
func (kc *KanjiCreate) AddRadicals(r ...*Radical) *KanjiCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return kc.AddRadicalIDs(ids...)
}

// Mutation returns the KanjiMutation object of the builder.
func (kc *KanjiCreate) Mutation() *KanjiMutation {
	return kc.mutation
}

// Save creates the Kanji in the database.
func (kc *KanjiCreate) Save(ctx context.Context) (*Kanji, error) {
	var (
		err  error
		node *Kanji
	)
	kc.defaults()
	if len(kc.hooks) == 0 {
		if err = kc.check(); err != nil {
			return nil, err
		}
		node, err = kc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KanjiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kc.check(); err != nil {
				return nil, err
			}
			kc.mutation = mutation
			if node, err = kc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kc.hooks) - 1; i >= 0; i-- {
			if kc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Kanji)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KanjiMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KanjiCreate) SaveX(ctx context.Context) *Kanji {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KanjiCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KanjiCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kc *KanjiCreate) defaults() {
	if _, ok := kc.mutation.CreatedAt(); !ok {
		v := kanji.DefaultCreatedAt()
		kc.mutation.SetCreatedAt(v)
	}
	if _, ok := kc.mutation.UpdatedAt(); !ok {
		v := kanji.DefaultUpdatedAt()
		kc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kc.mutation.ID(); !ok {
		v := kanji.DefaultID()
		kc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KanjiCreate) check() error {
	if _, ok := kc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Kanji.created_at"`)}
	}
	if _, ok := kc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Kanji.updated_at"`)}
	}
	if _, ok := kc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "Kanji.symbol"`)}
	}
	if v, ok := kc.mutation.Symbol(); ok {
		if err := kanji.SymbolValidator(v); err != nil {
			return &ValidationError{Name: "symbol", err: fmt.Errorf(`ent: validator failed for field "Kanji.symbol": %w`, err)}
		}
	}
	if _, ok := kc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Kanji.name"`)}
	}
	if v, ok := kc.mutation.Name(); ok {
		if err := kanji.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Kanji.name": %w`, err)}
		}
	}
	if _, ok := kc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Kanji.level"`)}
	}
	if v, ok := kc.mutation.Level(); ok {
		if err := kanji.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Kanji.level": %w`, err)}
		}
	}
	if _, ok := kc.mutation.Reading(); !ok {
		return &ValidationError{Name: "reading", err: errors.New(`ent: missing required field "Kanji.reading"`)}
	}
	if v, ok := kc.mutation.Reading(); ok {
		if err := kanji.ReadingValidator(v); err != nil {
			return &ValidationError{Name: "reading", err: fmt.Errorf(`ent: validator failed for field "Kanji.reading": %w`, err)}
		}
	}
	if _, ok := kc.mutation.Onyomi(); !ok {
		return &ValidationError{Name: "onyomi", err: errors.New(`ent: missing required field "Kanji.onyomi"`)}
	}
	if _, ok := kc.mutation.Kunyomi(); !ok {
		return &ValidationError{Name: "kunyomi", err: errors.New(`ent: missing required field "Kanji.kunyomi"`)}
	}
	if _, ok := kc.mutation.Nanori(); !ok {
		return &ValidationError{Name: "nanori", err: errors.New(`ent: missing required field "Kanji.nanori"`)}
	}
	if _, ok := kc.mutation.MeaningMnemonic(); !ok {
		return &ValidationError{Name: "meaning_mnemonic", err: errors.New(`ent: missing required field "Kanji.meaning_mnemonic"`)}
	}
	if v, ok := kc.mutation.MeaningMnemonic(); ok {
		if err := kanji.MeaningMnemonicValidator(v); err != nil {
			return &ValidationError{Name: "meaning_mnemonic", err: fmt.Errorf(`ent: validator failed for field "Kanji.meaning_mnemonic": %w`, err)}
		}
	}
	if _, ok := kc.mutation.ReadingMnemonic(); !ok {
		return &ValidationError{Name: "reading_mnemonic", err: errors.New(`ent: missing required field "Kanji.reading_mnemonic"`)}
	}
	if v, ok := kc.mutation.ReadingMnemonic(); ok {
		if err := kanji.ReadingMnemonicValidator(v); err != nil {
			return &ValidationError{Name: "reading_mnemonic", err: fmt.Errorf(`ent: validator failed for field "Kanji.reading_mnemonic": %w`, err)}
		}
	}
	return nil
}

func (kc *KanjiCreate) sqlSave(ctx context.Context) (*Kanji, error) {
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (kc *KanjiCreate) createSpec() (*Kanji, *sqlgraph.CreateSpec) {
	var (
		_node = &Kanji{config: kc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kanji.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kanji.FieldID,
			},
		}
	)
	if id, ok := kc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := kc.mutation.CreatedAt(); ok {
		_spec.SetField(kanji.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := kc.mutation.UpdatedAt(); ok {
		_spec.SetField(kanji.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := kc.mutation.Symbol(); ok {
		_spec.SetField(kanji.FieldSymbol, field.TypeString, value)
		_node.Symbol = value
	}
	if value, ok := kc.mutation.Name(); ok {
		_spec.SetField(kanji.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := kc.mutation.AltNames(); ok {
		_spec.SetField(kanji.FieldAltNames, field.TypeOther, value)
		_node.AltNames = value
	}
	if value, ok := kc.mutation.Similar(); ok {
		_spec.SetField(kanji.FieldSimilar, field.TypeOther, value)
		_node.Similar = value
	}
	if value, ok := kc.mutation.Level(); ok {
		_spec.SetField(kanji.FieldLevel, field.TypeInt32, value)
		_node.Level = value
	}
	if value, ok := kc.mutation.Reading(); ok {
		_spec.SetField(kanji.FieldReading, field.TypeString, value)
		_node.Reading = value
	}
	if value, ok := kc.mutation.Onyomi(); ok {
		_spec.SetField(kanji.FieldOnyomi, field.TypeOther, value)
		_node.Onyomi = value
	}
	if value, ok := kc.mutation.Kunyomi(); ok {
		_spec.SetField(kanji.FieldKunyomi, field.TypeOther, value)
		_node.Kunyomi = value
	}
	if value, ok := kc.mutation.Nanori(); ok {
		_spec.SetField(kanji.FieldNanori, field.TypeOther, value)
		_node.Nanori = value
	}
	if value, ok := kc.mutation.MeaningMnemonic(); ok {
		_spec.SetField(kanji.FieldMeaningMnemonic, field.TypeString, value)
		_node.MeaningMnemonic = value
	}
	if value, ok := kc.mutation.ReadingMnemonic(); ok {
		_spec.SetField(kanji.FieldReadingMnemonic, field.TypeString, value)
		_node.ReadingMnemonic = value
	}
	if nodes := kc.mutation.VocabulariesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   kanji.VocabulariesTable,
			Columns: kanji.VocabulariesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vocabulary.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kc.mutation.RadicalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   kanji.RadicalsTable,
			Columns: kanji.RadicalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: radical.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KanjiCreateBulk is the builder for creating many Kanji entities in bulk.
type KanjiCreateBulk struct {
	config
	builders []*KanjiCreate
}

// Save creates the Kanji entities in the database.
func (kcb *KanjiCreateBulk) Save(ctx context.Context) ([]*Kanji, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Kanji, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KanjiMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KanjiCreateBulk) SaveX(ctx context.Context) []*Kanji {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KanjiCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KanjiCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}