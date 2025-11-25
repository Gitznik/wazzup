package daemon

import (
	"context"
	"log"
	"time"

	probeservice "github.com/gitznik/wazzup/internal/service"
	"github.com/kardianos/service"
)

type Service struct {
	daemon *Daemon
	cancel context.CancelFunc
}

func (p *Service) Start(s service.Service) error {
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel

	serv := probeservice.NewFromDSN(probeservice.DefaultDSN)
	p.daemon = New(time.Minute, serv)

	go func() {
		if err := p.daemon.Start(ctx); err != nil {
			log.Printf("Daemon error: %v", err)
		}
	}()

	return nil
}

func (p *Service) Stop(s service.Service) error {
	if p.cancel != nil {
		p.cancel()
	}
	if err := p.daemon.Close(); err != nil {
		log.Printf("Failed shutting down daemon: %v", err)
	}
	return nil
}
