package main

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
)

func main() {
	plugin.Run(
		func(command.Cli) *cobra.Command {
			return &cobra.Command{
				Use: "sleep",
				RunE: func(_ *cobra.Command, _ []string) error {
					ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
					defer cancel()
					return exec.CommandContext(ctx, "sleep", "1000").Run()
				},
			}
		},
		manager.Metadata{
			SchemaVersion: "0.1.0",
			Vendor:        "Docker Inc.",
			Version:       "0.1.0",
		},
	)
}
