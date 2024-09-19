package calendar

import (
	"context"
	"errors"
	"log"
	"os"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const agenda = "Eventos"

var (
	ErrAGendaNotFound = errors.New("Agenda não encontrada")
)

type Calendar struct {
	Service *gCalendar.Service
}

func NewCLient() *Calendar {
	ctx := context.Background()
	credentials, err := os.ReadFile("./credentials.json")
	if err != nil {
		log.Fatalf("Erro ao ler credenciais: %v", err)
	}

	service, err := gCalendar.NewService(ctx, option.WithCredentialsJSON(credentials))
	if err != nil {
		log.Fatalf("Erro ao criar serviço: %v", err)
	}

	return &Calendar{
		Service: service,
	}
}

func (c *Calendar) GetAgendaId() (string, error) {
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Erro ao listar agendas: %v", err)
	}

	for _, v := range list.Items {
		if v.Summary == agenda {
			return v.Id, nil
		}
	}

	return "", ErrAGendaNotFound
}
