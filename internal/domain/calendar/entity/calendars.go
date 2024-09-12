package entity

type Calendars struct {
	Items []Calendar
}

type Calendar struct {
	Id          string
	Description string
}
