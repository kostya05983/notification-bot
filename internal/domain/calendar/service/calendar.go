package calendar

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

func GetAuthLink(config *oauth2.Config) string {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func GetToken(config *oauth2.Config, authCode string) (*oauth2.Token, error) {
	token, err := config.Exchange(context.TODO(), authCode)

	if err != nil {
		return nil, err
	}

	return token, nil
}

type CalendarService struct {
	api calendar.Service
}

func New(calendar calendar.Service) CalendarService {
	return CalendarService{calendar}
}

func (c *CalendarService) getEvents(ctx context.Context) error {
	t := time.Now().Format(time.RFC3339)

	events, err := c.api.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()

	if err != nil {
		return errors.Wrap(err, "c.api.Events")
	}

	return events
}
