package main

import (
	"github.com/gitznik/wazzup/internal/daemon"
	"github.com/spf13/cobra"

	"github.com/kardianos/service"
)

func serviceConfig() *service.Config {
	return &service.Config{
		Name:        "HealthProber",
		DisplayName: "Health Probing Daemon",
		Description: "Runs the health probes in the background",
	}
}

// installCmd represents the install command
func InstallCmd() *cobra.Command {
	return &cobra.Command{
		Use: "install",
		RunE: func(cmd *cobra.Command, args []string) error {
			svcConfig := serviceConfig()
			s, err := service.New(&daemon.Service{}, svcConfig)
			if err != nil {
				return err
			}
			return s.Install()
		},
	}
}

func UninstallCmd() *cobra.Command {
	return &cobra.Command{
		Use: "uninstall",
		RunE: func(cmd *cobra.Command, args []string) error {
			svcConfig := serviceConfig()
			s, err := service.New(&daemon.Service{}, svcConfig)
			if err != nil {
				return err
			}
			return s.Uninstall()
		},
	}
}

func init() {
	rootCmd.AddCommand(InstallCmd())
	rootCmd.AddCommand(UninstallCmd())
}
