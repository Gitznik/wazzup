package service

import (
	"github.com/gitznik/wazzup/ent"
	"github.com/gitznik/wazzup/internal/db"
)

type Service struct {
	Ent *ent.Client
}

const DefaultDSN = "file:wazzup.db?_fk=1"

func New() *Service {
	client := db.Startup(DefaultDSN)
	return &Service{
		Ent: client,
	}
}

func NewFromDSN(dbDSN string) *Service {
	client := db.Startup(dbDSN)
	return &Service{
		Ent: client,
	}
}

func (s Service) Close() error {
	return s.Ent.Close()
}
