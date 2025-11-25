/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
func CreateCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "create",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				return fmt.Errorf("name must not be an empty string")
			}
			url, _ := cmd.Flags().GetString("url")
			if url == "" {
				return fmt.Errorf("url must not be an empty string")
			}

			ctx, cancel := context.WithTimeout(cmd.Context(), 1*time.Second)
			defer cancel()
			s := ServiceFrom(cmd)
			p, err := s.CreateProbe(ctx, name, url)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Created probe: %+v", p)
			return nil
		},
	}

	cmd.Flags().StringP("name", "n", "", "Name of the probe")
	cmd.Flags().StringP("url", "u", "", "URL to monitor")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		panic(fmt.Sprintf("failed setting `name` flag as required: %v", err))
	}
	if err := cmd.MarkFlagRequired("url"); err != nil {
		panic(fmt.Sprintf("failed setting `url` flag as required: %v", err))
	}
	return &cmd
}

func init() {
	cmd := CreateCmd()
	rootCmd.AddCommand(cmd)
}
