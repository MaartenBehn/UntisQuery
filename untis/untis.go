package untis

import (
	"github.com/Stroby241/UntisAPI"
	"github.com/Stroby241/UntisQuerry/event"
	"github.com/Stroby241/UntisQuerry/state"
	"time"
)

var user *UntisAPI.User

func Init() {
	event.On(event.EventLogin, func(data interface{}) {
		strings := data.([4]string)
		login(strings[0], strings[1], strings[2], strings[3])
	})

	event.On(event.EventLogout, func(data interface{}) {
		logout()
	})

	event.On(event.EventAddTeacher, func(data interface{}) {
		strings := data.([2]string)
		success := addTeacher(strings[0], strings[1])
		if success {
			event.Go(event.EventUpdateQuerryPanel, nil)
			event.Go(event.EventSetPage, state.PageQuerry)
		}
	})

	event.On(event.EventLoadTimeTable, func(data interface{}) {
		loadTimetable(data.(time.Time))
	})

	event.On(event.EventQuerryTaecher, func(data interface{}) {
		queryTeacher(data.(*state.Teacher))
	})
}

func login(username string, password string, school string, server string) {
	user = UntisAPI.NewUser(username, password, school, server)
	err := user.Login()
	if err != nil {
		user = nil
		event.Go(event.EventLoginResult, false)
		return
	}

	go initCalls()
}

var rooms map[int]UntisAPI.Room
var classes map[int]UntisAPI.Class

func initCalls() {
	if user == nil {
		event.Go(event.EventLoginResult, false)
		return
	}

	event.Go(event.EventStartLoading, "Login")
	var err error
	rooms, err = user.GetRooms()
	if err != nil {
		event.Go(event.EventLoginResult, false)
		return
	}

	event.Go(event.EventUpdateLoading, 50.0)
	classes, err = user.GetClasses()
	if err != nil {
		event.Go(event.EventLoginResult, false)
		return
	}

	event.Go(event.EventLoginResult, true)
}

func logout() {
	if user == nil {
		return
	}
	user.Logout()
	user = nil

	event.Go(event.EventSetPage, state.PageStart)
}

var timetable []map[int]UntisAPI.Period
var date int

func loadTimetable(dateTime time.Time) bool {
	newDate := UntisAPI.ToUntisDate(dateTime)

	if date == newDate {
		return true
	}

	date = newDate

	timetable = []map[int]UntisAPI.Period{}
	counter := 0

	event.Go(event.EventStartLoading, "Timetable")
	for _, room := range rooms {
		event.Go(event.EventUpdateLoading, float64(counter)/float64(len(rooms))*100.0)

		periods, err := user.GetTimeTable(room.Id, 4, date, date)
		if err != nil {
			return false
		}

		if periods != nil {
			timetable = append(timetable, periods)
		}

		counter++
	}
	event.Go(event.EventSetPage, state.PageQuerry)

	return true
}
