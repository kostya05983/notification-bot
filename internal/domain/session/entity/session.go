package entity

import (
	"fmt"
)

type Session struct {
	ChatId      int64
	State       State
	GoogleToken *string
}

type State int64

var StateCreated State = 1
var StateWaitToken State = 2
var StateChooseCalendar State = 3

func (s *Session) AddToken(token string) error {
	if s.State != StateWaitToken {
		return fmt.Errorf("State is wrong to update token %d", s.State)
	}

	s.GoogleToken = &token

	return nil
}

func (s *Session) IsAuthenticated() bool {
	return s.GoogleToken != nil
}
