package state

import "github.com/Stroby241/UntisAPI"

var Teachers []*Teacher

type Teacher struct {
	Firstname string
	Lastname  string
	Id        int
}

const (
	PageStart      = 0
	PageLogin      = 1
	PageQuerry     = 2
	PageAddTeacher = 3
	PageLoading    = 4
	PageMax        = 5
)

var FoundPeriods []UntisAPI.Period
