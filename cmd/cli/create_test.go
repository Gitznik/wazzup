package main

import (
	"fmt"
	"testing"

	"github.com/gitznik/wazzup/internal/service"
	"github.com/gitznik/wazzup/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// Given a database client and service are available
	client, _ := testhelpers.DBClient(t)
	s := service.Service{Ent: client}
	defer func() { _ = client.Close() }()
	defer func() { _ = s.Close() }()

	t.Run("create", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			// Given valid probe parameters
			cmd := CreateCmd()

			// When executing the create command with valid arguments
			r, err := callCLI(t, cmd, &s, []string{"--name", t.Name(), "--url", "https://example.com"})

			// Then the command should succeed and output probe details
			assert.NoError(t, err)
			assert.Contains(t, string(r), t.Name())
			assert.Contains(t, string(r), "example.com")
		})

		t.Run("short_flags", func(t *testing.T) {
			// Given valid probe parameters using short flags
			cmd := CreateCmd()

			// When executing the create command with short flags
			r, err := callCLI(t, cmd, &s, []string{"-n", t.Name(), "-u", "https://short.example.com"})

			// Then the command should succeed
			assert.NoError(t, err)
			assert.Contains(t, string(r), t.Name())
			assert.Contains(t, string(r), "short.example.com")
		})

		t.Run("missing_name_flag", func(t *testing.T) {
			// Given only URL is provided without name
			cmd := CreateCmd()

			// When executing the create command without name flag
			_, err := callCLI(t, cmd, &s, []string{"--url", "https://example.com"})

			// Then the command should fail due to missing required flag
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "name")
		})

		t.Run("missing_url_flag", func(t *testing.T) {
			// Given only name is provided without URL
			cmd := CreateCmd()

			// When executing the create command without URL flag
			_, err := callCLI(t, cmd, &s, []string{"--name", t.Name()})

			// Then the command should fail due to missing required flag
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "url")
		})

		t.Run("empty_name_value", func(t *testing.T) {
			// Given empty name value is provided
			cmd := CreateCmd()

			// When executing the create command with empty name
			_, err := callCLI(t, cmd, &s, []string{"--name", "", "--url", "https://example.com"})

			// Then the command should fail with validation error
			assert.Error(t, err)
		})

		t.Run("empty_url_value", func(t *testing.T) {
			// Given empty URL value is provided
			cmd := CreateCmd()

			// When executing the create command with empty URL
			_, err := callCLI(t, cmd, &s, []string{"--name", t.Name(), "--url", ""})

			// Then the command should fail with validation error
			assert.Error(t, err)
		})

		t.Run("special_characters_in_name", func(t *testing.T) {
			// Given probe names with various special characters
			testCases := []struct {
				nameSuffix string
			}{
				{"-with-dashes"},
				{"_with_underscores"},
				{".with.dots"},
				{" with spaces"},
				{"@with#symbols"},
			}

			for _, tc := range testCases {
				t.Run("name"+tc.nameSuffix, func(t *testing.T) {
					cmd := CreateCmd()
					name := t.Name() + tc.nameSuffix

					// When creating probe with special characters in name
					r, err := callCLI(t, cmd, &s, []string{"--name", name, "--url", "https://example.com"})

					// Then command should succeed and handle special characters
					assert.NoError(t, err)
					assert.Contains(t, string(r), name)
				})
			}
		})

		t.Run("various_url_formats", func(t *testing.T) {
			// Given different URL formats
			testCases := []struct {
				url         string
				description string
			}{
				{"https://example.com", "standard_https"},
				{"http://example.com", "standard_http"},
				{"https://subdomain.example.com", "subdomain"},
				{"https://example.com:8080", "with_port"},
				{"https://example.com/path", "with_path"},
				{"https://example.com/path?query=value", "with_query"},
			}

			for i, tc := range testCases {
				t.Run("url_"+tc.description, func(t *testing.T) {
					cmd := CreateCmd()
					name := t.Name() + fmt.Sprintf("-%d", i)

					// When creating probe with different URL formats
					r, err := callCLI(t, cmd, &s, []string{"--name", name, "--url", tc.url})

					// Then command should succeed for valid URLs
					assert.NoError(t, err)
					assert.Contains(t, string(r), "Created probe:")
				})
			}
		})

		t.Run("duplicate_names", func(t *testing.T) {
			// Given a probe with a specific name already exists
			cmd1 := CreateCmd()
			probeName := t.Name() + "-duplicate"

			// When creating the first probe
			r1, err1 := callCLI(t, cmd1, &s, []string{"--name", probeName, "--url", "https://first.example.com"})
			assert.NoError(t, err1)
			assert.Contains(t, string(r1), probeName)

			// And when attempting to create another probe with the same name
			cmd2 := CreateCmd()
			_, err2 := callCLI(t, cmd2, &s, []string{"--name", probeName, "--url", "https://second.example.com"})

			// Then the second creation should fail due to uniqueness constraint
			assert.Error(t, err2)
		})

		t.Run("long_name", func(t *testing.T) {
			// Given a very long probe name
			longName := t.Name() + "-very-long-name-that-exceeds-normal-limits-and-tests-boundary-conditions"
			cmd := CreateCmd()

			// When creating probe with long name
			r, err := callCLI(t, cmd, &s, []string{"--name", longName, "--url", "https://example.com"})

			// Then system should handle long names appropriately
			assert.NoError(t, err)
			assert.Contains(t, string(r), "Created probe:")
		})

		t.Run("multiple_sequential", func(t *testing.T) {
			// Given multiple probe creation requests
			probes := []struct {
				suffix string
				url    string
			}{
				{"-seq1", "https://seq1.example.com"},
				{"-seq2", "https://seq2.example.com"},
				{"-seq3", "https://seq3.example.com"},
			}

			// When creating multiple probes sequentially
			for i, probe := range probes {
				cmd := CreateCmd()
				name := t.Name() + probe.suffix
				r, err := callCLI(t, cmd, &s, []string{"--name", name, "--url", probe.url})

				// Then each probe should be created successfully
				assert.NoError(t, err, "probe %d should be created successfully", i+1)
				assert.Contains(t, string(r), name)
				assert.Contains(t, string(r), "Created probe:")
			}
		})
	})
}
