package service

import (
	"context"

	"github.com/gitznik/wazzup/ent"
	"github.com/gitznik/wazzup/internal/repository"
)

func (s Service) GetResults(ctx context.Context, name string) ([]*ent.HealthProbeResults, error) {
	return repository.GetHealthProbeResultsForName(ctx, s.Ent, name)
}

func (s Service) GetProbes(ctx context.Context) (ent.HealthProbes, error) {
	return repository.GetProbes(ctx, s.Ent)
}
