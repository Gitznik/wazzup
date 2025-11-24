package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-playground/validator/v10"
)

// HealthProbe holds the schema definition for the HealthProbe entity.
type HealthProbe struct {
	ent.Schema
}

// Fields of the HealthProbe.
func (HealthProbe) Fields() []ent.Field {
	validate := validator.New()
	return []ent.Field{
		field.String("name").Unique(),
		field.String("url").Validate(func(s string) error {
			return validate.Var(s, "required,url")
		}),
	}
}

// Edges of the HealthProbe.
func (HealthProbe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("health_probe_results", HealthProbeResults.Type),
	}
}
