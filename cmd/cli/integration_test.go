package main

import (
	"testing"

	"github.com/gitznik/wazzup/internal/service"
	"github.com/gitznik/wazzup/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	client := testhelpers.DBClient(t)
	s := service.Service{Ent: client}
	defer func() { _ = client.Close() }()
	defer func() { _ = s.Close() }()

	probeName := t.Name()

	t.Run("add probe", func(t *testing.T) {
		cmd := CreateCmd()

		r, err := callCLI(t, cmd, &s, []string{"--name", probeName, "--url", "https://example.com"})
		t.Logf("Got result %s", r)
		assert.NoError(t, err)
	})

	t.Run("add result", func(t *testing.T) {
		cmd := CreateProbeResultCmd()

		_, err := callCLI(t, cmd, &s, []string{"1", "1"})
		assert.NoError(t, err)
	})

	t.Run("get results", func(t *testing.T) {
		cmd := GetCmd()

		_, err := callCLI(t, cmd, &s, []string{probeName})
		assert.NoError(t, err)
	})
}
