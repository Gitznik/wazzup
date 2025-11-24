package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Probe holds the schema definition for the Probe entity.
type Probe struct {
	ent.Schema
}

// Fields of the Probe.
func (Probe) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Probe.
func (Probe) Edges() []ent.Edge {
	return nil
}
