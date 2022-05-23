// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/predicate"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/todoitem"
)

// TodoItemUpdate is the builder for updating TodoItem entities.
type TodoItemUpdate struct {
	config
	hooks    []Hook
	mutation *TodoItemMutation
}

// Where appends a list predicates to the TodoItemUpdate builder.
func (tiu *TodoItemUpdate) Where(ps ...predicate.TodoItem) *TodoItemUpdate {
	tiu.mutation.Where(ps...)
	return tiu
}

// SetTitle sets the "title" field.
func (tiu *TodoItemUpdate) SetTitle(s string) *TodoItemUpdate {
	tiu.mutation.SetTitle(s)
	return tiu
}

// SetCompleted sets the "completed" field.
func (tiu *TodoItemUpdate) SetCompleted(b bool) *TodoItemUpdate {
	tiu.mutation.SetCompleted(b)
	return tiu
}

// SetOrder sets the "order" field.
func (tiu *TodoItemUpdate) SetOrder(i int) *TodoItemUpdate {
	tiu.mutation.ResetOrder()
	tiu.mutation.SetOrder(i)
	return tiu
}

// AddOrder adds i to the "order" field.
func (tiu *TodoItemUpdate) AddOrder(i int) *TodoItemUpdate {
	tiu.mutation.AddOrder(i)
	return tiu
}

// SetCreatedAt sets the "created_at" field.
func (tiu *TodoItemUpdate) SetCreatedAt(t time.Time) *TodoItemUpdate {
	tiu.mutation.SetCreatedAt(t)
	return tiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tiu *TodoItemUpdate) SetNillableCreatedAt(t *time.Time) *TodoItemUpdate {
	if t != nil {
		tiu.SetCreatedAt(*t)
	}
	return tiu
}

// SetUpdatedAt sets the "updated_at" field.
func (tiu *TodoItemUpdate) SetUpdatedAt(t time.Time) *TodoItemUpdate {
	tiu.mutation.SetUpdatedAt(t)
	return tiu
}

// Mutation returns the TodoItemMutation object of the builder.
func (tiu *TodoItemUpdate) Mutation() *TodoItemMutation {
	return tiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tiu *TodoItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tiu.defaults()
	if len(tiu.hooks) == 0 {
		affected, err = tiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tiu.mutation = mutation
			affected, err = tiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tiu.hooks) - 1; i >= 0; i-- {
			if tiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tiu *TodoItemUpdate) SaveX(ctx context.Context) int {
	affected, err := tiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tiu *TodoItemUpdate) Exec(ctx context.Context) error {
	_, err := tiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tiu *TodoItemUpdate) ExecX(ctx context.Context) {
	if err := tiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tiu *TodoItemUpdate) defaults() {
	if _, ok := tiu.mutation.UpdatedAt(); !ok {
		v := todoitem.UpdateDefaultUpdatedAt()
		tiu.mutation.SetUpdatedAt(v)
	}
}

func (tiu *TodoItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   todoitem.Table,
			Columns: todoitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todoitem.FieldID,
			},
		},
	}
	if ps := tiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tiu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todoitem.FieldTitle,
		})
	}
	if value, ok := tiu.mutation.Completed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: todoitem.FieldCompleted,
		})
	}
	if value, ok := tiu.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todoitem.FieldOrder,
		})
	}
	if value, ok := tiu.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todoitem.FieldOrder,
		})
	}
	if value, ok := tiu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todoitem.FieldCreatedAt,
		})
	}
	if value, ok := tiu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todoitem.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todoitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TodoItemUpdateOne is the builder for updating a single TodoItem entity.
type TodoItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoItemMutation
}

// SetTitle sets the "title" field.
func (tiuo *TodoItemUpdateOne) SetTitle(s string) *TodoItemUpdateOne {
	tiuo.mutation.SetTitle(s)
	return tiuo
}

// SetCompleted sets the "completed" field.
func (tiuo *TodoItemUpdateOne) SetCompleted(b bool) *TodoItemUpdateOne {
	tiuo.mutation.SetCompleted(b)
	return tiuo
}

// SetOrder sets the "order" field.
func (tiuo *TodoItemUpdateOne) SetOrder(i int) *TodoItemUpdateOne {
	tiuo.mutation.ResetOrder()
	tiuo.mutation.SetOrder(i)
	return tiuo
}

// AddOrder adds i to the "order" field.
func (tiuo *TodoItemUpdateOne) AddOrder(i int) *TodoItemUpdateOne {
	tiuo.mutation.AddOrder(i)
	return tiuo
}

// SetCreatedAt sets the "created_at" field.
func (tiuo *TodoItemUpdateOne) SetCreatedAt(t time.Time) *TodoItemUpdateOne {
	tiuo.mutation.SetCreatedAt(t)
	return tiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tiuo *TodoItemUpdateOne) SetNillableCreatedAt(t *time.Time) *TodoItemUpdateOne {
	if t != nil {
		tiuo.SetCreatedAt(*t)
	}
	return tiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tiuo *TodoItemUpdateOne) SetUpdatedAt(t time.Time) *TodoItemUpdateOne {
	tiuo.mutation.SetUpdatedAt(t)
	return tiuo
}

// Mutation returns the TodoItemMutation object of the builder.
func (tiuo *TodoItemUpdateOne) Mutation() *TodoItemMutation {
	return tiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tiuo *TodoItemUpdateOne) Select(field string, fields ...string) *TodoItemUpdateOne {
	tiuo.fields = append([]string{field}, fields...)
	return tiuo
}

// Save executes the query and returns the updated TodoItem entity.
func (tiuo *TodoItemUpdateOne) Save(ctx context.Context) (*TodoItem, error) {
	var (
		err  error
		node *TodoItem
	)
	tiuo.defaults()
	if len(tiuo.hooks) == 0 {
		node, err = tiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tiuo.mutation = mutation
			node, err = tiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tiuo.hooks) - 1; i >= 0; i-- {
			if tiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tiuo *TodoItemUpdateOne) SaveX(ctx context.Context) *TodoItem {
	node, err := tiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tiuo *TodoItemUpdateOne) Exec(ctx context.Context) error {
	_, err := tiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tiuo *TodoItemUpdateOne) ExecX(ctx context.Context) {
	if err := tiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tiuo *TodoItemUpdateOne) defaults() {
	if _, ok := tiuo.mutation.UpdatedAt(); !ok {
		v := todoitem.UpdateDefaultUpdatedAt()
		tiuo.mutation.SetUpdatedAt(v)
	}
}

func (tiuo *TodoItemUpdateOne) sqlSave(ctx context.Context) (_node *TodoItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   todoitem.Table,
			Columns: todoitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todoitem.FieldID,
			},
		},
	}
	id, ok := tiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TodoItem.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todoitem.FieldID)
		for _, f := range fields {
			if !todoitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todoitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tiuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todoitem.FieldTitle,
		})
	}
	if value, ok := tiuo.mutation.Completed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: todoitem.FieldCompleted,
		})
	}
	if value, ok := tiuo.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todoitem.FieldOrder,
		})
	}
	if value, ok := tiuo.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: todoitem.FieldOrder,
		})
	}
	if value, ok := tiuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todoitem.FieldCreatedAt,
		})
	}
	if value, ok := tiuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todoitem.FieldUpdatedAt,
		})
	}
	_node = &TodoItem{config: tiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todoitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
