package main

import (
	"context"

	"github.com/gitznik/wazzup/internal/service"
	"github.com/spf13/cobra"
)

type serviceKey struct{}

func WithService(cmd *cobra.Command, service *service.Service) {
	parent := cmd.Context()
	if parent == nil {
		parent = context.Background()
	}
	ctx := context.WithValue(parent, serviceKey{}, service)
	cmd.SetContext(ctx)
}

func ServiceFrom(cmd *cobra.Command) *service.Service {
	db, _ := cmd.Context().Value(serviceKey{}).(*service.Service)
	return db
}
