package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"
)
// Title holds the schema definition for the Title entity.
type Title struct {
	ent.Schema
}

// Fields of the Title.
func (Title) Fields() []ent.Field {
	return []ent.Field{
		field.String("Title_type").NotEmpty(),
	}
}

// Edges of the Title.
func (Title) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("patients", Patient.Type).StorageKey(edge.Column("title_id")),
    }
}
