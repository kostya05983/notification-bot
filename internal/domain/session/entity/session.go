package entity

import (
	"fmt"
)

type Session struct {
	ChatId      int64
	State       State
	GoogleToken *string
	CalendarId  *string
}

type State int64

var StateCreated State = 1
var StateWaitToken State = 2
var StateChooseCalendar State = 3
var StateWaitCalendar State = 4
var StateActive State = 5

func (s *Session) AddToken(token string) error {
	if s.State != StateWaitToken {
		return fmt.Errorf("State is wrong to update token %d", s.State)
	}

	s.GoogleToken = &token
	s.State = StateChooseCalendar

	return nil
}

func (s *Session) IsAuthenticated() bool {
	return s.GoogleToken != nil
}

func (s *Session) WaitToken() error {
	if s.State != StateCreated {
		return fmt.Errorf("State is wrong to update token %d", s.State)
	}

	s.State = StateWaitToken

	return nil
}

func (s *Session) WaitCalendar() error {
	if s.State != StateChooseCalendar {
		return fmt.Errorf("State is wrong to update token %d", s.State)
	}

	s.State = StateWaitCalendar

	return nil
}

func (s *Session) AddCalendar(calendarId string) error {
	if s.State != StateWaitCalendar {
		return fmt.Errorf("State is wrong to update token %d", s.State)
	}

	s.CalendarId = &calendarId
	s.State = StateActive

	return nil
}
