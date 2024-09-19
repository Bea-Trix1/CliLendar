package cmd

import (
	"fmt"
	"os"

	cmd "github.com/Bea-Trix1/CliLendar/cmd/events"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "calendar",
		Short: "Gerencie sua agenda",
		// Args:  cobra.ExactArgs(2),
	}

	rootCmd.AddCommand(cmd.EventsCmd)

	return &rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
