package entity

type Session struct {
	ChatId      int64
	State       State
	GoogleToken *string
}

type State string

var Created State
var WaitToken State
var ChooseCalendar State
