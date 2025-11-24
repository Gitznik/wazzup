package repository

import (
	"context"
	"fmt"

	"github.com/gitznik/wazzup/ent"
	"github.com/gitznik/wazzup/ent/healthprobe"
	"github.com/gitznik/wazzup/ent/schema"
)

func CreateHealthProbe(ctx context.Context, client *ent.Client, name string, url string) (*ent.HealthProbe, error) {
	p, err := client.HealthProbe.Create().SetName(name).SetURL(url).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed storing health probe: %w", err)
	}
	return p, nil
}

func GetHealthProbeResultsForId(ctx context.Context, client *ent.Client, probe int) ([]*ent.HealthProbeResults, error) {
	p, err := client.HealthProbe.Get(ctx, probe)
	if err != nil {
		return nil, fmt.Errorf("failed getting health probe: %w", err)
	}
	r, err := p.QueryHealthProbeResults().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting health probe results: %w", err)
	}
	return r, err
}

func GetHealthProbeResultsForName(ctx context.Context, client *ent.Client, probe string) ([]*ent.HealthProbeResults, error) {
	p, err := client.HealthProbe.Query().Where(healthprobe.Name(probe)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting health probe: %w", err)
	}
	r, err := p.QueryHealthProbeResults().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting health probe results: %w", err)
	}
	return r, err
}

func CreateHealthProbeResults(ctx context.Context, client *ent.Client, probe int, result schema.CheckResult) (*ent.HealthProbeResults, error) {
	p, err := client.HealthProbeResults.Create().SetResult(result).SetHealthProbeID(probe).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed storing health probe result: %w", err)
	}
	return p, nil
}
