package entity

type Session struct {
	chatId      int64
	state       State
	googleToken *string
}

type State string

var Created State
var WaitToken State
var ChooseCalendar State
