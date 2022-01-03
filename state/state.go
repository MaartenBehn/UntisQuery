package state

import "time"

var Teachers []*Teacher

type Teacher struct {
	Firstname string
	Lastname  string
	Id        int
}

var Periods []*Period

type Period struct {
	StartTime time.Time
	EndTime   time.Time
	Classes   []string
	Subjects  []string
	Rooms     []string
}
