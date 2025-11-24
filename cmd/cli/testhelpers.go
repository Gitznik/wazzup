package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/gitznik/wazzup/internal/service"
	"github.com/spf13/cobra"
)

func callCLI(t *testing.T, cmd *cobra.Command, service *service.Service, args []string) ([]byte, error) {
	t.Helper()
	WithService(cmd, service)
	cmd.SetArgs(args)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	if err := cmd.Execute(); err != nil {
		return nil, err
	}
	return io.ReadAll(b)
}
