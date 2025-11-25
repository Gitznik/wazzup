package daemon

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gitznik/wazzup/internal/service"
)

type Daemon struct {
	interval time.Duration
	task     Prober
	service  *service.Service
}

func (d *Daemon) Close() error {
	return d.service.Close()
}

type Prober = func(context.Context, *http.Client, *service.Service) error

func New(d time.Duration, s *service.Service) *Daemon {
	return &Daemon{
		interval: d,
		service:  s,
		task:     probe,
	}
}

func (d *Daemon) Start(ctx context.Context) error {
	fmt.Println("Starting Daemon")
	ticker := time.NewTicker(d.interval)
	defer ticker.Stop()

	client := http.Client{Timeout: 5 * time.Second}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	if err := d.task(ctx, &client, d.service); err != nil {
		return err
	}

	for {
		select {
		case <-ticker.C:
			if err := d.task(ctx, &client, d.service); err != nil {
				return err
			}
		case <-sigChan:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
