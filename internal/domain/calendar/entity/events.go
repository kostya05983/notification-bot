package entity

import (
	"time"
)

type Events struct {
	items    []Event
	timeZone string
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
	start             time.Time
	originalStartTime time.Time
}
