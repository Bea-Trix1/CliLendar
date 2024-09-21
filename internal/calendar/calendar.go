package calendar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	gCalendar "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const agenda = "Eventos"

var (
	ErrAGendaNotFound = errors.New("agenda não encontrada")
	ErrAddAgenda      = errors.New("erro ao adicionar agenda")
	ErrListEvents     = errors.New("erro ao listar eventos")
)

type Calendar struct {
	Service    *gCalendar.Service
	CalendarId string
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

func (c *Calendar) GetAgendaId() error {
	list, err := c.Service.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Erro ao listar agendas: %v", err)
	}

	for _, v := range list.Items {
		if v.Summary == agenda {
			c.CalendarId = v.Id
		}
	}

	return nil
}

func (c *Calendar) InsertAgenda(id string) error {
	entry := &gCalendar.CalendarListEntry{
		Id: c.CalendarId,
	}

	_, err := c.Service.CalendarList.Insert(entry).Do()
	if err != nil {
		return ErrAddAgenda
	}

	return nil
}

func (c *Calendar) ListEventsWeek() error {
	now := time.Now()
	weekday := now.Weekday()
	startDate := now.AddDate(0, 0, -int(weekday))
	endDate := startDate.AddDate(0, 0, 7)

	events, err := c.Service.Events.List(c.CalendarId).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return ErrListEvents
	}

	for _, v := range events.Items {
		fmt.Printf("Nome do evento: %s | status: %s | quando: %s \n", v.Summary, v.Status, v.Start.DateTime)
	}

	return nil
}

func (c *Calendar) ListEventsToday() error {
	year, month, day := time.Now().Date()
	startDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 0, 1)

	events, err := c.Service.Events.List(c.CalendarId).TimeMin(startDate.Format(time.RFC3339)).TimeMax(endDate.Format(time.RFC3339)).Do()
	if err != nil {
		return ErrListEvents
	}

	if len(events.Items) == 0 {
		fmt.Println("Nenhum evento encontrado para hoje.")
		return nil
	}

	for _, v := range events.Items {
		fmt.Printf("Nome do evento: %s | status: %s | quando: %s \n", v.Summary, v.Status, v.Start.DateTime)
	}

	return nil
}
