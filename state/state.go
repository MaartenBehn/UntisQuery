package state

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