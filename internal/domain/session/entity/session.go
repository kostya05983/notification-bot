package entity

type Session struct {
	ChatId      int64
	State       State
	GoogleToken *string
}

type State int64

var Created State = 1
var WaitToken State = 2
var ChooseCalendar State = 3
