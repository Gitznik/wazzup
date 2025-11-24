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
		Use:   "create",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			url, _ := cmd.Flags().GetString("url")

			ctx, cancel := context.WithTimeout(cmd.Context(), 1*time.Second)
			defer cancel()
			s := ServiceFrom(cmd)
			defer func() {
				if err := s.Close(); err != nil {
					fmt.Fprintf(cmd.OutOrStdout(), "Failed closing the db: %v\n", err)
				}
			}()
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
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("url")
	return &cmd
}

func init() {
	cmd := CreateCmd()
	rootCmd.AddCommand(cmd)
}
