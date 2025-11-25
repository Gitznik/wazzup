package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CheckResult int

const (
	Success CheckResult = iota
	Failure
	HTTPFailure
)

func (c CheckResult) IsValid() bool {
	return c >= Success && c <= Failure
}

// HealthProbeResults holds the schema definition for the HealthProbeResults entity.
type HealthProbeResults struct {
	ent.Schema
}

// Fields of the HealthProbeResults.
func (HealthProbeResults) Fields() []ent.Field {
	return []ent.Field{
		field.Int("result").GoType(CheckResult(0)).NonNegative(),
		field.String("context").Optional(),
	}
}

// Edges of the HealthProbeResults.
func (HealthProbeResults) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("health_probe", HealthProbe.Type).Ref("health_probe_results").Unique(),
	}
}
