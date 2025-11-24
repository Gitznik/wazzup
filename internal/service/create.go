package service

import (
	"context"

	"github.com/gitznik/wazzup/ent"
	"github.com/gitznik/wazzup/ent/schema"
	"github.com/gitznik/wazzup/internal/repository"
)

func (s Service) CreateProbe(ctx context.Context, name, url string) (*ent.HealthProbe, error) {
	return repository.CreateHealthProbe(ctx, s.Ent, name, url)
}

func (s Service) CreateProbeResult(ctx context.Context, id int, result schema.CheckResult) (*ent.HealthProbeResults, error) {
	return repository.CreateHealthProbeResults(ctx, s.Ent, id, result)
}
