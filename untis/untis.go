package untis

import (
	"errors"
	"github.com/Stroby241/UntisAPI"
	"github.com/Stroby241/UntisQuery/event"
	"github.com/Stroby241/UntisQuery/state"
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
		addTeacher(strings[0], strings[1])
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
		event.Go(event.EventLoginResult, err)
		return
	}

	go initCalls()
}

var rooms map[int]UntisAPI.Room
var classes map[int]UntisAPI.Class
var subjects map[int]UntisAPI.Subject

func initCalls() {
	event.Go(event.EventLoading, 0.0)
	if user == nil {
		event.Go(event.EventLoginResult, errors.New("No User"))
		return
	}

	event.Go(event.EventLoading, 0.1)
	var err error
	rooms, err = user.GetRooms()
	if err != nil {
		event.Go(event.EventLoginResult, err)
		return
	}

	event.Go(event.EventLoading, 0.5)
	classes, err = user.GetClasses()
	if err != nil {
		event.Go(event.EventLoginResult, err)
		return
	}

	event.Go(event.EventLoading, 0.8)
	subjects, err = user.GetSubjects()
	if err != nil {
		event.Go(event.EventLoginResult, err)
		return
	}

	event.Go(event.EventLoading, 1.0)
	event.Go(event.EventLoginResult, nil)
}

func logout() {
	if user == nil {
		return
	}
	user.Logout()
	user = nil
}

var timetable []map[int]UntisAPI.Period
var currentTimetable []UntisAPI.Period
var fromDate int
var toDate int

const timetablePredicting = 7

func loadTimetable(dateTime time.Time) {
	untisDate := UntisAPI.ToUntisDate(dateTime)

	if !(untisDate >= fromDate && untisDate <= toDate) {
		fromDate = UntisAPI.ToUntisDate(dateTime.AddDate(0, 0, -timetablePredicting))
		toDate = UntisAPI.ToUntisDate(dateTime.AddDate(0, 0, timetablePredicting))

		timetable = []map[int]UntisAPI.Period{}
		counter := 0
		event.Go(event.EventLoading, 0.0)
		for _, room := range rooms {
			event.Go(event.EventLoading, float64(counter)/float64(len(rooms)))
			periods, err := user.GetTimeTable(room.Id, 4, fromDate, toDate)
			if err != nil {
				event.Go(event.EventLoadTimeTableResult, err)
				return
			}

			if periods != nil {
				timetable = append(timetable, periods)
			}

			counter++
		}
		event.Go(event.EventLoading, 1.0)
	}

	currentTimetable = []UntisAPI.Period{}
	for _, room := range timetable {
		for _, period := range room {
			if period.Date == untisDate {
				currentTimetable = append(currentTimetable, period)
			}
		}
	}

	event.Go(event.EventLoadTimeTableResult, nil)
}
