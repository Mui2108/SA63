// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/panupong/app/ent/job"
	"github.com/panupong/app/ent/patient"
	"github.com/panupong/app/ent/predicate"
)

// JobUpdate is the builder for updating Job entities.
type JobUpdate struct {
	config
	hooks      []Hook
	mutation   *JobMutation
	predicates []predicate.Job
}

// Where adds a new predicate for the builder.
func (ju *JobUpdate) Where(ps ...predicate.Job) *JobUpdate {
	ju.predicates = append(ju.predicates, ps...)
	return ju
}

// SetJobName sets the Job_name field.
func (ju *JobUpdate) SetJobName(s string) *JobUpdate {
	ju.mutation.SetJobName(s)
	return ju
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (ju *JobUpdate) AddPatientIDs(ids ...int) *JobUpdate {
	ju.mutation.AddPatientIDs(ids...)
	return ju
}

// AddPatients adds the patients edges to Patient.
func (ju *JobUpdate) AddPatients(p ...*Patient) *JobUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ju.AddPatientIDs(ids...)
}

// Mutation returns the JobMutation object of the builder.
func (ju *JobUpdate) Mutation() *JobMutation {
	return ju.mutation
}

// RemovePatientIDs removes the patients edge to Patient by ids.
func (ju *JobUpdate) RemovePatientIDs(ids ...int) *JobUpdate {
	ju.mutation.RemovePatientIDs(ids...)
	return ju
}

// RemovePatients removes patients edges to Patient.
func (ju *JobUpdate) RemovePatients(p ...*Patient) *JobUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ju.RemovePatientIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ju *JobUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ju.mutation.JobName(); ok {
		if err := job.JobNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Job_name", err: fmt.Errorf("ent: validator failed for field \"Job_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ju.hooks) == 0 {
		affected, err = ju.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ju.mutation = mutation
			affected, err = ju.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ju.hooks) - 1; i >= 0; i-- {
			mut = ju.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ju.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ju *JobUpdate) SaveX(ctx context.Context) int {
	affected, err := ju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ju *JobUpdate) Exec(ctx context.Context) error {
	_, err := ju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ju *JobUpdate) ExecX(ctx context.Context) {
	if err := ju.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ju *JobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		},
	}
	if ps := ju.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ju.mutation.JobName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldJobName,
		})
	}
	if nodes := ju.mutation.RemovedPatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.PatientsTable,
			Columns: []string{job.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.PatientsTable,
			Columns: []string{job.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// JobUpdateOne is the builder for updating a single Job entity.
type JobUpdateOne struct {
	config
	hooks    []Hook
	mutation *JobMutation
}

// SetJobName sets the Job_name field.
func (juo *JobUpdateOne) SetJobName(s string) *JobUpdateOne {
	juo.mutation.SetJobName(s)
	return juo
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (juo *JobUpdateOne) AddPatientIDs(ids ...int) *JobUpdateOne {
	juo.mutation.AddPatientIDs(ids...)
	return juo
}

// AddPatients adds the patients edges to Patient.
func (juo *JobUpdateOne) AddPatients(p ...*Patient) *JobUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return juo.AddPatientIDs(ids...)
}

// Mutation returns the JobMutation object of the builder.
func (juo *JobUpdateOne) Mutation() *JobMutation {
	return juo.mutation
}

// RemovePatientIDs removes the patients edge to Patient by ids.
func (juo *JobUpdateOne) RemovePatientIDs(ids ...int) *JobUpdateOne {
	juo.mutation.RemovePatientIDs(ids...)
	return juo
}

// RemovePatients removes patients edges to Patient.
func (juo *JobUpdateOne) RemovePatients(p ...*Patient) *JobUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return juo.RemovePatientIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (juo *JobUpdateOne) Save(ctx context.Context) (*Job, error) {
	if v, ok := juo.mutation.JobName(); ok {
		if err := job.JobNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Job_name", err: fmt.Errorf("ent: validator failed for field \"Job_name\": %w", err)}
		}
	}

	var (
		err  error
		node *Job
	)
	if len(juo.hooks) == 0 {
		node, err = juo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			juo.mutation = mutation
			node, err = juo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(juo.hooks) - 1; i >= 0; i-- {
			mut = juo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, juo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (juo *JobUpdateOne) SaveX(ctx context.Context) *Job {
	j, err := juo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return j
}

// Exec executes the query on the entity.
func (juo *JobUpdateOne) Exec(ctx context.Context) error {
	_, err := juo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (juo *JobUpdateOne) ExecX(ctx context.Context) {
	if err := juo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (juo *JobUpdateOne) sqlSave(ctx context.Context) (j *Job, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		},
	}
	id, ok := juo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Job.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := juo.mutation.JobName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldJobName,
		})
	}
	if nodes := juo.mutation.RemovedPatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.PatientsTable,
			Columns: []string{job.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.PatientsTable,
			Columns: []string{job.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	j = &Job{config: juo.config}
	_spec.Assign = j.assignValues
	_spec.ScanValues = j.scanValues()
	if err = sqlgraph.UpdateNode(ctx, juo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return j, nil
}