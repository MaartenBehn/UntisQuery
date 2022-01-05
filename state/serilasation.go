package state

import (
	"bytes"
	"encoding/gob"
	"github.com/Stroby241/UntisQuery/event"
	"os"
)

func getPath() string {
	if android {
		return os.Getenv("TMPDIR") + "/"
	}
	return "data/"
}

func SaveLogins() {
	path := getPath()

	var buffer bytes.Buffer
	e := gob.NewEncoder(&buffer)

	if err := e.Encode(Logins); err != nil {
		event.Go(event.EventHandleError, err)
	}

	saveBufferToFile(path+"login.txt", &buffer)
}

func LoadLogins() {
	path := getPath()

	buffer := loadBufferFromFile(path + "login.txt")
	if buffer == nil {
		return
	}

	t := &[]*Login{}
	d := gob.NewDecoder(buffer)
	event.Go(event.EventHandleError, d.Decode(t))

	Logins = *t
}

func SaveTeacher() {
	path := getPath()

	var buffer bytes.Buffer
	e := gob.NewEncoder(&buffer)

	if err := e.Encode(Teachers); err != nil {
		event.Go(event.EventHandleError, err)
	}

	saveBufferToFile(path+"teacher.txt", &buffer)
}

func LoadTeacher() {
	path := getPath()

	buffer := loadBufferFromFile(path + "teacher.txt")
	if buffer == nil {
		return
	}

	t := &[]*Teacher{}
	d := gob.NewDecoder(buffer)
	event.Go(event.EventHandleError, d.Decode(t))

	Teachers = *t
}

func saveBufferToFile(path string, buffer *bytes.Buffer) {
	f, err := os.Create(path)
	event.Go(event.EventHandleError, err)
	if err != nil {
		return
	}

	_, err = f.Write(buffer.Bytes())
	event.Go(event.EventHandleError, err)

	defer event.Go(event.EventHandleError, f.Close())
}

func loadBufferFromFile(path string) *bytes.Buffer {
	buf, err := os.ReadFile(path)
	event.Go(event.EventHandleError, err)

	return bytes.NewBuffer(buf)
}
