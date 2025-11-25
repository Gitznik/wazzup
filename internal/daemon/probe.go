package daemon

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gitznik/wazzup/ent/schema"
	"github.com/gitznik/wazzup/internal/service"
)

func probe(ctx context.Context, client *http.Client, service *service.Service) error {
	fmt.Println("Running probe")
	probes, err := service.GetProbes(ctx)
	if err != nil {
		fmt.Printf("Got error %s\n", err)
		return err
	}
	fmt.Printf("Got probes: %+v\n", probes)
	for _, p := range probes {
		fmt.Printf("Probe to %s\n", p.URL)
		var s schema.CheckResult
		var c string
		res, err := client.Get(p.URL)
		if err != nil {
			s = schema.Failure
			c = err.Error()
		} else {
			defer func() { _ = res.Body.Close() }()
			if res.StatusCode > 299 {
				s = schema.HTTPFailure
				c = fmt.Sprintf("Failed with status %d", res.StatusCode)
			} else {
				s = schema.Success
			}
		}
		fmt.Printf("Probe result of probe %d: status %d - result %d, context %s\n", p.ID, res.StatusCode, s, c)
		if _, err := service.CreateProbeResult(ctx, p.ID, s, c); err != nil {
			fmt.Printf("Failed storing result: %s\n", err)
			return err
		}
	}
	fmt.Println("Finished loop")
	return nil
}
