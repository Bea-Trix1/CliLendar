package cmd

import (
	"fmt"
	"log"

	calendar "github.com/Bea-Trix1/CliLendar/internal/calendar"
	"github.com/spf13/cobra"
)

func init() {
	EventsCmd.AddCommand(EventsWeekCmd)
	EventsCmd.AddCommand(EventsTodayCmd)
}

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Listar eventos",
	Long:  "Listar todos os eventos de sua agenda",
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewCLient()
		err := c.GetAgendaId()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Sucesso ao buscar eventos!")
	},
}
