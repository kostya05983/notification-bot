package entity

import (
	"time"
)

type Events struct {
	Items    []Event
	TimeZone string
}

type EventStatus string

var Confirmed EventStatus = "confirmed"
var Tenative EventStatus = "tentative"
var Cancelled EventStatus = "cancelled"

type Event struct {
	Id                string
	Summary           string
	Description       string
	Status            EventStatus
	Start             time.Time
	OriginalStartTime time.Time
}
