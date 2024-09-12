package calendar

import (
	"context"
	"notification-bot/internal/domain/calendar/entity"
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

func (c *CalendarService) getEvents(ctx context.Context) (*entity.Events, error) {
	t := time.Now().Format(time.RFC3339)

	events, err := c.api.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()

	if err != nil {
		return nil, errors.Wrap(err, "c.api.Events")
	}

	entity, err := toEntity(events)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert to entity")
	}

	return entity, nil
}

func toEntity(src *calendar.Events) (*entity.Events, error) {
	result := make([]entity.Event, 0, len(src.Items))

	for _, item := range src.Items {
		startTime, err := parseTimeRfc(item.Start.DateTime)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse start datetime")
		}

		originalStartTime, err := parseTimeRfc(item.OriginalStartTime.DateTime)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse original start datetime")
		}

		result = append(result, entity.Event{
			Id:                item.Id,
			Summary:           item.Summary,
			Description:       item.Description,
			Status:            entity.EventStatus(item.Status),
			Start:             startTime,
			OriginalStartTime: originalStartTime,
		})
	}

	return &entity.Events{
		Items:    result,
		TimeZone: src.TimeZone,
	}, nil
}

func parseTimeRfc(value string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	t, err := time.Parse(layout, str)

	return t, err
}
