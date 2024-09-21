package cmd

import (
	"log"

	calendar "github.com/Bea-Trix1/CliLendar/internal/calendar"
	"github.com/spf13/cobra"
)

var EventsWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "Listar eventos da semana",
	Long:  "Listar todos os eventos da semana",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewCLient()
		err := c.GetAgendaId()
		if err != nil {
			log.Fatal(err.Error())
		}
		c.ListEventsWeek()
	},
}
