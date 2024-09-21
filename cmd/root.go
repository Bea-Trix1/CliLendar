package cmd

import (
	"fmt"
	"os"

	agenda "github.com/Bea-Trix1/CliLendar/cmd/agenda"
	cmd "github.com/Bea-Trix1/CliLendar/cmd/events"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "calendario",
		Short: "Gerencie sua agenda",
	}

	rootCmd.AddCommand(cmd.EventsCmd, agenda.AgendaCmd)

	return &rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
