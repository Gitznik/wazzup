package service

import (
	"github.com/gitznik/wazzup/ent"
	"github.com/gitznik/wazzup/internal/db"
)

type Service struct {
	Ent *ent.Client
}

func New() *Service {
	client := db.Startup()
	return &Service{
		Ent: client,
	}
}

func (s Service) Close() error {
	return s.Ent.Close()
}
