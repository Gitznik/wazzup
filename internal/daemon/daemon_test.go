package daemon_test

import (
	"context"
	"testing"
	"time"

	"github.com/gitznik/wazzup/internal/daemon"
	"github.com/gitznik/wazzup/internal/service"
	"github.com/gitznik/wazzup/internal/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDaemon(t *testing.T) {
	client := testhelpers.DBClient(t)
	defer func() { _ = client.Close() }()
	s := service.Service{Ent: client}
	daemon := daemon.New(time.Millisecond, &s)
	defer func() {
		if err := daemon.Close(); err != nil {
			t.Fatalf("failed closing daemon: %s", err)
		}
	}()

	t.Run("start probe", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		p, err := client.HealthProbe.Create().SetName(t.Name()).SetURL("https://example.com").Save(ctx)
		require.NoError(t, err)
		probeDL := 1 * time.Second
		dctx, cancel := context.WithTimeout(ctx, probeDL)
		defer cancel()
		err = daemon.Start(dctx)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
		r, err := p.QueryHealthProbeResults().All(ctx)
		assert.NoError(t, err)
		assert.Less(t, 0, len(r))
	})
}
