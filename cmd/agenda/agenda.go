package cmd

import (
	"fmt"
	"log"

	calendar "github.com/Bea-Trix1/CliLendar/internal/calendar"
	"github.com/spf13/cobra"
)

var AgendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "Listar agendas",
	Long:  "Listar todas suas agendas",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := calendar.NewCLient()
		err := c.InsertAgenda(args[0])
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Sucesso ao buscar agenda!")
	},
}
