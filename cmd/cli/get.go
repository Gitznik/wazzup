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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <name>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		ctx, cancel := context.WithTimeout(cmd.Context(), 1*time.Second)
		defer cancel()
		s := ServiceFrom(cmd)
		defer func() {
			if err := s.Close(); err != nil {
				fmt.Printf("Failed closing the db: %v\n", err)
			}
		}()
		p, err := s.GetResults(ctx, name)
		if err != nil {
			return err
		}
		fmt.Printf("Results: %+v", p)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
