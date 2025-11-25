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

func GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get <name>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			ctx, cancel := context.WithTimeout(cmd.Context(), 1*time.Second)
			defer cancel()
			s := ServiceFrom(cmd)
			p, err := s.GetResults(ctx, name)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Results: %+v", p)
			return nil
		}}
	return cmd

}

func init() {
	cmd := GetCmd()
	rootCmd.AddCommand(cmd)
}
