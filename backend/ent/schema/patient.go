package schema

import (
	
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("Card_id").NotEmpty(),
		field.String("First_name").NotEmpty(),
		field.String("Last_name").NotEmpty(),
		field.String("Allergic").NotEmpty(),
		field.String("Address").NotEmpty(),
		field.Int("Age"),
		field.Time("Birthday"),
		
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gender", Gender.Type).
			Ref("patients").
			Unique(),
		edge.From("title", Title.Type).
			Ref("patients").
			Unique(),
		edge.From("job", Job.Type).
			Ref("patients").
			Unique(),
	}
}
