/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gitznik/wazzup/ent/schema"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
func CreateProbeResultCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "create-probe-result <probe-id> <result>",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(cmd.Context(), 1*time.Second)
			defer cancel()
			pid, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("failed converting project id to integer: %w", err)
			}
			r, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("failed converting result to integer: %w", err)
			}
			rp := schema.CheckResult(r)
			if !rp.IsValid() {
				return fmt.Errorf("passed result %d is out of range", r)
			}
			s := ServiceFrom(cmd)
			p, err := s.CreateProbeResult(ctx, pid, rp, "")
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Created probe: %+v", p)
			return nil
		},
	}
	return &cmd
}

func init() {
	cmd := CreateProbeResultCmd()
	rootCmd.AddCommand(cmd)
}
