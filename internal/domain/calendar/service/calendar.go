package service

import (
	"notification-bot/internal/domain/calendar/entity"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/calendar/v3"
)

type CalendarService struct {
	api calendar.Service
}

func New(calendar calendar.Service) CalendarService {
	return CalendarService{calendar}
}

func (c *CalendarService) GetEvents(lastTime *time.Time) (*entity.Events, error) {
	t := time.Now().Format(time.RFC3339)

	request := c.api.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime")

	if lastTime != nil {
		request = request.TimeMax(lastTime.Format("2011-06-03T10:00:00-07:00"))
	}

	events, err := request.Do()

	if err != nil {
		return nil, errors.Wrap(err, "c.api.Events")
	}

	entity, err := toEntityEvents(events)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert to entity")
	}

	return entity, nil
}

func (c *CalendarService) GetCalendars() (*entity.Calendars, error) {
	list, err := c.api.CalendarList.List().Do()

	if err != nil {
		return nil, errors.Wrap(err, "c.api.CalendarList")
	}

	return toEntityCalendars(list), nil
}

func toEntityCalendars(src *calendar.CalendarList) *entity.Calendars {
	calendars := make([]entity.Calendar, 0, len(src.Items))

	for _, item := range src.Items {
		calendars = append(calendars, entity.Calendar{
			Description: item.Description,
			Id:          item.Id,
		})
	}

	return &entity.Calendars{
		Items: calendars,
	}
}

func toEntityEvents(src *calendar.Events) (*entity.Events, error) {
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
