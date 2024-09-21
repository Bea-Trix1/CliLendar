package cmd

import (
	"log"

	calendar "github.com/Bea-Trix1/CliLendar/internal/calendar"
	"github.com/spf13/cobra"
)

var EventsTodayCmd = &cobra.Command{
	Use:   "today",
	Short: "Listar eventos de hoje",
	Long:  "Listar todos os eventos de hoje",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewCLient()
		err := c.GetAgendaId()
		if err != nil {
			log.Fatal(err.Error())
		}
		c.ListEventsToday()
	},
}
