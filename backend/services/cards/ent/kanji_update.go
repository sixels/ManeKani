// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"sixels.io/manekani/services/cards/ent/kanji"
	"sixels.io/manekani/services/cards/ent/predicate"
	"sixels.io/manekani/services/cards/ent/radical"
	"sixels.io/manekani/services/cards/ent/vocabulary"
)

// KanjiUpdate is the builder for updating Kanji entities.
type KanjiUpdate struct {
	config
	hooks    []Hook
	mutation *KanjiMutation
}

// Where appends a list predicates to the KanjiUpdate builder.
func (ku *KanjiUpdate) Where(ps ...predicate.Kanji) *KanjiUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetUpdatedAt sets the "updated_at" field.
func (ku *KanjiUpdate) SetUpdatedAt(t time.Time) *KanjiUpdate {
	ku.mutation.SetUpdatedAt(t)
	return ku
}

// SetName sets the "name" field.
func (ku *KanjiUpdate) SetName(s string) *KanjiUpdate {
	ku.mutation.SetName(s)
	return ku
}

// SetAltNames sets the "alt_names" field.
func (ku *KanjiUpdate) SetAltNames(pa pgtype.TextArray) *KanjiUpdate {
	ku.mutation.SetAltNames(pa)
	return ku
}

// SetNillableAltNames sets the "alt_names" field if the given value is not nil.
func (ku *KanjiUpdate) SetNillableAltNames(pa *pgtype.TextArray) *KanjiUpdate {
	if pa != nil {
		ku.SetAltNames(*pa)
	}
	return ku
}

// ClearAltNames clears the value of the "alt_names" field.
func (ku *KanjiUpdate) ClearAltNames() *KanjiUpdate {
	ku.mutation.ClearAltNames()
	return ku
}

// SetLevel sets the "level" field.
func (ku *KanjiUpdate) SetLevel(i int32) *KanjiUpdate {
	ku.mutation.ResetLevel()
	ku.mutation.SetLevel(i)
	return ku
}

// AddLevel adds i to the "level" field.
func (ku *KanjiUpdate) AddLevel(i int32) *KanjiUpdate {
	ku.mutation.AddLevel(i)
	return ku
}

// SetReading sets the "reading" field.
func (ku *KanjiUpdate) SetReading(s string) *KanjiUpdate {
	ku.mutation.SetReading(s)
	return ku
}

// SetOnyomi sets the "onyomi" field.
func (ku *KanjiUpdate) SetOnyomi(pa pgtype.TextArray) *KanjiUpdate {
	ku.mutation.SetOnyomi(pa)
	return ku
}

// SetKunyomi sets the "kunyomi" field.
func (ku *KanjiUpdate) SetKunyomi(pa pgtype.TextArray) *KanjiUpdate {
	ku.mutation.SetKunyomi(pa)
	return ku
}

// SetNanori sets the "nanori" field.
func (ku *KanjiUpdate) SetNanori(pa pgtype.TextArray) *KanjiUpdate {
	ku.mutation.SetNanori(pa)
	return ku
}

// SetMeaningMnemonic sets the "meaning_mnemonic" field.
func (ku *KanjiUpdate) SetMeaningMnemonic(s string) *KanjiUpdate {
	ku.mutation.SetMeaningMnemonic(s)
	return ku
}

// SetReadingMnemonic sets the "reading_mnemonic" field.
func (ku *KanjiUpdate) SetReadingMnemonic(s string) *KanjiUpdate {
	ku.mutation.SetReadingMnemonic(s)
	return ku
}

// AddVocabularyIDs adds the "vocabularies" edge to the Vocabulary entity by IDs.
func (ku *KanjiUpdate) AddVocabularyIDs(ids ...uuid.UUID) *KanjiUpdate {
	ku.mutation.AddVocabularyIDs(ids...)
	return ku
}

// AddVocabularies adds the "vocabularies" edges to the Vocabulary entity.
func (ku *KanjiUpdate) AddVocabularies(v ...*Vocabulary) *KanjiUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ku.AddVocabularyIDs(ids...)
}

// AddRadicalIDs adds the "radicals" edge to the Radical entity by IDs.
func (ku *KanjiUpdate) AddRadicalIDs(ids ...uuid.UUID) *KanjiUpdate {
	ku.mutation.AddRadicalIDs(ids...)
	return ku
}

// AddRadicals adds the "radicals" edges to the Radical entity.
func (ku *KanjiUpdate) AddRadicals(r ...*Radical) *KanjiUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ku.AddRadicalIDs(ids...)
}

// Mutation returns the KanjiMutation object of the builder.
func (ku *KanjiUpdate) Mutation() *KanjiMutation {
	return ku.mutation
}

// ClearVocabularies clears all "vocabularies" edges to the Vocabulary entity.
func (ku *KanjiUpdate) ClearVocabularies() *KanjiUpdate {
	ku.mutation.ClearVocabularies()
	return ku
}

// RemoveVocabularyIDs removes the "vocabularies" edge to Vocabulary entities by IDs.
func (ku *KanjiUpdate) RemoveVocabularyIDs(ids ...uuid.UUID) *KanjiUpdate {
	ku.mutation.RemoveVocabularyIDs(ids...)
	return ku
}

// RemoveVocabularies removes "vocabularies" edges to Vocabulary entities.
func (ku *KanjiUpdate) RemoveVocabularies(v ...*Vocabulary) *KanjiUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ku.RemoveVocabularyIDs(ids...)
}

// ClearRadicals clears all "radicals" edges to the Radical entity.
func (ku *KanjiUpdate) ClearRadicals() *KanjiUpdate {
	ku.mutation.ClearRadicals()
	return ku
}

// RemoveRadicalIDs removes the "radicals" edge to Radical entities by IDs.
func (ku *KanjiUpdate) RemoveRadicalIDs(ids ...uuid.UUID) *KanjiUpdate {
	ku.mutation.RemoveRadicalIDs(ids...)
	return ku
}

// RemoveRadicals removes "radicals" edges to Radical entities.
func (ku *KanjiUpdate) RemoveRadicals(r ...*Radical) *KanjiUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ku.RemoveRadicalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KanjiUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ku.defaults()
	if len(ku.hooks) == 0 {
		if err = ku.check(); err != nil {
			return 0, err
		}
		affected, err = ku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KanjiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ku.check(); err != nil {
				return 0, err
			}
			ku.mutation = mutation
			affected, err = ku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ku.hooks) - 1; i >= 0; i-- {
			if ku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KanjiUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KanjiUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KanjiUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ku *KanjiUpdate) defaults() {
	if _, ok := ku.mutation.UpdatedAt(); !ok {
		v := kanji.UpdateDefaultUpdatedAt()
		ku.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ku *KanjiUpdate) check() error {
	if v, ok := ku.mutation.Name(); ok {
		if err := kanji.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Kanji.name": %w`, err)}
		}
	}
	if v, ok := ku.mutation.Level(); ok {
		if err := kanji.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Kanji.level": %w`, err)}
		}
	}
	if v, ok := ku.mutation.Reading(); ok {
		if err := kanji.ReadingValidator(v); err != nil {
			return &ValidationError{Name: "reading", err: fmt.Errorf(`ent: validator failed for field "Kanji.reading": %w`, err)}
		}
	}
	if v, ok := ku.mutation.MeaningMnemonic(); ok {
		if err := kanji.MeaningMnemonicValidator(v); err != nil {
			return &ValidationError{Name: "meaning_mnemonic", err: fmt.Errorf(`ent: validator failed for field "Kanji.meaning_mnemonic": %w`, err)}
		}
	}
	if v, ok := ku.mutation.ReadingMnemonic(); ok {
		if err := kanji.ReadingMnemonicValidator(v); err != nil {
			return &ValidationError{Name: "reading_mnemonic", err: fmt.Errorf(`ent: validator failed for field "Kanji.reading_mnemonic": %w`, err)}
		}
	}
	return nil
}

func (ku *KanjiUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kanji.Table,
			Columns: kanji.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kanji.FieldID,
			},
		},
	}
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.UpdatedAt(); ok {
		_spec.SetField(kanji.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ku.mutation.Name(); ok {
		_spec.SetField(kanji.FieldName, field.TypeString, value)
	}
	if value, ok := ku.mutation.AltNames(); ok {
		_spec.SetField(kanji.FieldAltNames, field.TypeOther, value)
	}
	if ku.mutation.AltNamesCleared() {
		_spec.ClearField(kanji.FieldAltNames, field.TypeOther)
	}
	if value, ok := ku.mutation.Level(); ok {
		_spec.SetField(kanji.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := ku.mutation.AddedLevel(); ok {
		_spec.AddField(kanji.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := ku.mutation.Reading(); ok {
		_spec.SetField(kanji.FieldReading, field.TypeString, value)
	}
	if value, ok := ku.mutation.Onyomi(); ok {
		_spec.SetField(kanji.FieldOnyomi, field.TypeOther, value)
	}
	if value, ok := ku.mutation.Kunyomi(); ok {
		_spec.SetField(kanji.FieldKunyomi, field.TypeOther, value)
	}
	if value, ok := ku.mutation.Nanori(); ok {
		_spec.SetField(kanji.FieldNanori, field.TypeOther, value)
	}
	if value, ok := ku.mutation.MeaningMnemonic(); ok {
		_spec.SetField(kanji.FieldMeaningMnemonic, field.TypeString, value)
	}
	if value, ok := ku.mutation.ReadingMnemonic(); ok {
		_spec.SetField(kanji.FieldReadingMnemonic, field.TypeString, value)
	}
	if ku.mutation.VocabulariesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ku.mutation.RemovedVocabulariesIDs(); len(nodes) > 0 && !ku.mutation.VocabulariesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ku.mutation.VocabulariesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ku.mutation.RadicalsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ku.mutation.RemovedRadicalsIDs(); len(nodes) > 0 && !ku.mutation.RadicalsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ku.mutation.RadicalsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kanji.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// KanjiUpdateOne is the builder for updating a single Kanji entity.
type KanjiUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KanjiMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (kuo *KanjiUpdateOne) SetUpdatedAt(t time.Time) *KanjiUpdateOne {
	kuo.mutation.SetUpdatedAt(t)
	return kuo
}

// SetName sets the "name" field.
func (kuo *KanjiUpdateOne) SetName(s string) *KanjiUpdateOne {
	kuo.mutation.SetName(s)
	return kuo
}

// SetAltNames sets the "alt_names" field.
func (kuo *KanjiUpdateOne) SetAltNames(pa pgtype.TextArray) *KanjiUpdateOne {
	kuo.mutation.SetAltNames(pa)
	return kuo
}

// SetNillableAltNames sets the "alt_names" field if the given value is not nil.
func (kuo *KanjiUpdateOne) SetNillableAltNames(pa *pgtype.TextArray) *KanjiUpdateOne {
	if pa != nil {
		kuo.SetAltNames(*pa)
	}
	return kuo
}

// ClearAltNames clears the value of the "alt_names" field.
func (kuo *KanjiUpdateOne) ClearAltNames() *KanjiUpdateOne {
	kuo.mutation.ClearAltNames()
	return kuo
}

// SetLevel sets the "level" field.
func (kuo *KanjiUpdateOne) SetLevel(i int32) *KanjiUpdateOne {
	kuo.mutation.ResetLevel()
	kuo.mutation.SetLevel(i)
	return kuo
}

// AddLevel adds i to the "level" field.
func (kuo *KanjiUpdateOne) AddLevel(i int32) *KanjiUpdateOne {
	kuo.mutation.AddLevel(i)
	return kuo
}

// SetReading sets the "reading" field.
func (kuo *KanjiUpdateOne) SetReading(s string) *KanjiUpdateOne {
	kuo.mutation.SetReading(s)
	return kuo
}

// SetOnyomi sets the "onyomi" field.
func (kuo *KanjiUpdateOne) SetOnyomi(pa pgtype.TextArray) *KanjiUpdateOne {
	kuo.mutation.SetOnyomi(pa)
	return kuo
}

// SetKunyomi sets the "kunyomi" field.
func (kuo *KanjiUpdateOne) SetKunyomi(pa pgtype.TextArray) *KanjiUpdateOne {
	kuo.mutation.SetKunyomi(pa)
	return kuo
}

// SetNanori sets the "nanori" field.
func (kuo *KanjiUpdateOne) SetNanori(pa pgtype.TextArray) *KanjiUpdateOne {
	kuo.mutation.SetNanori(pa)
	return kuo
}

// SetMeaningMnemonic sets the "meaning_mnemonic" field.
func (kuo *KanjiUpdateOne) SetMeaningMnemonic(s string) *KanjiUpdateOne {
	kuo.mutation.SetMeaningMnemonic(s)
	return kuo
}

// SetReadingMnemonic sets the "reading_mnemonic" field.
func (kuo *KanjiUpdateOne) SetReadingMnemonic(s string) *KanjiUpdateOne {
	kuo.mutation.SetReadingMnemonic(s)
	return kuo
}

// AddVocabularyIDs adds the "vocabularies" edge to the Vocabulary entity by IDs.
func (kuo *KanjiUpdateOne) AddVocabularyIDs(ids ...uuid.UUID) *KanjiUpdateOne {
	kuo.mutation.AddVocabularyIDs(ids...)
	return kuo
}

// AddVocabularies adds the "vocabularies" edges to the Vocabulary entity.
func (kuo *KanjiUpdateOne) AddVocabularies(v ...*Vocabulary) *KanjiUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return kuo.AddVocabularyIDs(ids...)
}

// AddRadicalIDs adds the "radicals" edge to the Radical entity by IDs.
func (kuo *KanjiUpdateOne) AddRadicalIDs(ids ...uuid.UUID) *KanjiUpdateOne {
	kuo.mutation.AddRadicalIDs(ids...)
	return kuo
}

// AddRadicals adds the "radicals" edges to the Radical entity.
func (kuo *KanjiUpdateOne) AddRadicals(r ...*Radical) *KanjiUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return kuo.AddRadicalIDs(ids...)
}

// Mutation returns the KanjiMutation object of the builder.
func (kuo *KanjiUpdateOne) Mutation() *KanjiMutation {
	return kuo.mutation
}

// ClearVocabularies clears all "vocabularies" edges to the Vocabulary entity.
func (kuo *KanjiUpdateOne) ClearVocabularies() *KanjiUpdateOne {
	kuo.mutation.ClearVocabularies()
	return kuo
}

// RemoveVocabularyIDs removes the "vocabularies" edge to Vocabulary entities by IDs.
func (kuo *KanjiUpdateOne) RemoveVocabularyIDs(ids ...uuid.UUID) *KanjiUpdateOne {
	kuo.mutation.RemoveVocabularyIDs(ids...)
	return kuo
}

// RemoveVocabularies removes "vocabularies" edges to Vocabulary entities.
func (kuo *KanjiUpdateOne) RemoveVocabularies(v ...*Vocabulary) *KanjiUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return kuo.RemoveVocabularyIDs(ids...)
}

// ClearRadicals clears all "radicals" edges to the Radical entity.
func (kuo *KanjiUpdateOne) ClearRadicals() *KanjiUpdateOne {
	kuo.mutation.ClearRadicals()
	return kuo
}

// RemoveRadicalIDs removes the "radicals" edge to Radical entities by IDs.
func (kuo *KanjiUpdateOne) RemoveRadicalIDs(ids ...uuid.UUID) *KanjiUpdateOne {
	kuo.mutation.RemoveRadicalIDs(ids...)
	return kuo
}

// RemoveRadicals removes "radicals" edges to Radical entities.
func (kuo *KanjiUpdateOne) RemoveRadicals(r ...*Radical) *KanjiUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return kuo.RemoveRadicalIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KanjiUpdateOne) Select(field string, fields ...string) *KanjiUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated Kanji entity.
func (kuo *KanjiUpdateOne) Save(ctx context.Context) (*Kanji, error) {
	var (
		err  error
		node *Kanji
	)
	kuo.defaults()
	if len(kuo.hooks) == 0 {
		if err = kuo.check(); err != nil {
			return nil, err
		}
		node, err = kuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KanjiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kuo.check(); err != nil {
				return nil, err
			}
			kuo.mutation = mutation
			node, err = kuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kuo.hooks) - 1; i >= 0; i-- {
			if kuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (kuo *KanjiUpdateOne) SaveX(ctx context.Context) *Kanji {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KanjiUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KanjiUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kuo *KanjiUpdateOne) defaults() {
	if _, ok := kuo.mutation.UpdatedAt(); !ok {
		v := kanji.UpdateDefaultUpdatedAt()
		kuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kuo *KanjiUpdateOne) check() error {
	if v, ok := kuo.mutation.Name(); ok {
		if err := kanji.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Kanji.name": %w`, err)}
		}
	}
	if v, ok := kuo.mutation.Level(); ok {
		if err := kanji.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Kanji.level": %w`, err)}
		}
	}
	if v, ok := kuo.mutation.Reading(); ok {
		if err := kanji.ReadingValidator(v); err != nil {
			return &ValidationError{Name: "reading", err: fmt.Errorf(`ent: validator failed for field "Kanji.reading": %w`, err)}
		}
	}
	if v, ok := kuo.mutation.MeaningMnemonic(); ok {
		if err := kanji.MeaningMnemonicValidator(v); err != nil {
			return &ValidationError{Name: "meaning_mnemonic", err: fmt.Errorf(`ent: validator failed for field "Kanji.meaning_mnemonic": %w`, err)}
		}
	}
	if v, ok := kuo.mutation.ReadingMnemonic(); ok {
		if err := kanji.ReadingMnemonicValidator(v); err != nil {
			return &ValidationError{Name: "reading_mnemonic", err: fmt.Errorf(`ent: validator failed for field "Kanji.reading_mnemonic": %w`, err)}
		}
	}
	return nil
}

func (kuo *KanjiUpdateOne) sqlSave(ctx context.Context) (_node *Kanji, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kanji.Table,
			Columns: kanji.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kanji.FieldID,
			},
		},
	}
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Kanji.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kanji.FieldID)
		for _, f := range fields {
			if !kanji.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != kanji.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.UpdatedAt(); ok {
		_spec.SetField(kanji.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := kuo.mutation.Name(); ok {
		_spec.SetField(kanji.FieldName, field.TypeString, value)
	}
	if value, ok := kuo.mutation.AltNames(); ok {
		_spec.SetField(kanji.FieldAltNames, field.TypeOther, value)
	}
	if kuo.mutation.AltNamesCleared() {
		_spec.ClearField(kanji.FieldAltNames, field.TypeOther)
	}
	if value, ok := kuo.mutation.Level(); ok {
		_spec.SetField(kanji.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := kuo.mutation.AddedLevel(); ok {
		_spec.AddField(kanji.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := kuo.mutation.Reading(); ok {
		_spec.SetField(kanji.FieldReading, field.TypeString, value)
	}
	if value, ok := kuo.mutation.Onyomi(); ok {
		_spec.SetField(kanji.FieldOnyomi, field.TypeOther, value)
	}
	if value, ok := kuo.mutation.Kunyomi(); ok {
		_spec.SetField(kanji.FieldKunyomi, field.TypeOther, value)
	}
	if value, ok := kuo.mutation.Nanori(); ok {
		_spec.SetField(kanji.FieldNanori, field.TypeOther, value)
	}
	if value, ok := kuo.mutation.MeaningMnemonic(); ok {
		_spec.SetField(kanji.FieldMeaningMnemonic, field.TypeString, value)
	}
	if value, ok := kuo.mutation.ReadingMnemonic(); ok {
		_spec.SetField(kanji.FieldReadingMnemonic, field.TypeString, value)
	}
	if kuo.mutation.VocabulariesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuo.mutation.RemovedVocabulariesIDs(); len(nodes) > 0 && !kuo.mutation.VocabulariesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuo.mutation.VocabulariesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if kuo.mutation.RadicalsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuo.mutation.RemovedRadicalsIDs(); len(nodes) > 0 && !kuo.mutation.RadicalsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuo.mutation.RadicalsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Kanji{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kanji.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
