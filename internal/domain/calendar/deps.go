package calendar

import (
	"notification-bot/internal/domain/calendar/entity"
	"time"

	"golang.org/x/oauth2"
)

type CalendarService interface {
	GetAuthLink(config *oauth2.Config) string
	GetToken(config *oauth2.Config, authCode string) (*oauth2.Token, error)
	GetCalendars() (*entity.Calendars, error)
	GetEvents(lastTime *time.Time) (*entity.Events, error)
}
