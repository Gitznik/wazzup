package repository

import (
	"context"
	"log"

	"github.com/gitznik/wazzup/ent"
)

func CreateProbe(ctx context.Context, client *ent.Client) (*ent.Probe, error) {
	p, err := client.Probe.Create().SetName("test").Save(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("probe was created: ", p)
	return p, nil
}
