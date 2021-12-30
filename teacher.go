package main

var teachers []*teacher

type teacher struct {
	firstname string
	lastname  string
	id        int // -1 not queried, -2 not valid
}

func (t *teacher) getId() {
	id, err := user.GetPersonId(t.firstname, t.lastname, true)
	if err != nil {
		t.id = -2
		return
	}
	t.id = id
}
