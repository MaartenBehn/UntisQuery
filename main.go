package main

import (
	"UntisQuerry/UntisV2"
	"fmt"
	"time"
)

var user *UntisV2.User
var timetable []map[int]UntisV2.Period
var rooms map[int]UntisV2.Room
var classes map[int]UntisV2.Class

func main() {
	user = UntisV2.NewUser("maarten8", "behn500", "TBZ Mitte Bremen", "https://tipo.webuntis.com")
	user.Login()

	date := UntisV2.ToUntisDate(time.Now())
	loadTimetable(date, date)

	querryTeacher("Daniel", "Dibbern")

	user.Logout()
}

func loadTimetable(startDate int, endDate int) {
	classes = user.GetClasses()
	rooms = user.GetRooms()

	timetable = []map[int]UntisV2.Period{}
	counter := 0
	for _, room := range rooms {
		fmt.Printf("Loading timetable of room: %d of %d. \r", counter, len(rooms))

		periods := user.GetTimeTable(room.Id, 4, startDate, endDate)

		if periods != nil {
			timetable = append(timetable, periods)
		}

		counter++
	}
}

func querryTeacher(firstname string, lastname string) {
	id := user.GetPersonId(firstname, lastname, true)

	var foundPeriod []UntisV2.Period
	for _, periods := range timetable {
		for _, period := range periods {
			for _, teacher := range period.Teacher {
				if teacher == id {
					foundPeriod = append(foundPeriod, period)
				}
			}
		}
	}

	fmt.Printf("%s %s found in %d Periods\n", firstname, lastname, len(foundPeriod))
	for _, period := range foundPeriod {
		fmt.Print("Room: ")
		for _, roomId := range period.Rooms {
			fmt.Printf("%s ", rooms[roomId].Name)
		}

		fmt.Print("Class: ")
		for _, classId := range period.Classes {
			fmt.Printf("%s ", classes[classId].Name)
		}

		date := UntisV2.ToGoDate(period.Date)
		fromTime := UntisV2.ToGoTime(period.StartTime)
		tillTime := UntisV2.ToGoTime(period.EndTime)

		fmt.Printf("Date: %d %s %d From: %02d:%02d  Till: %02d:%02d ",
			date.Day(), date.Month(), date.Year(),
			fromTime.Hour(), fromTime.Minute(),
			tillTime.Hour(), tillTime.Minute())

		fmt.Print("\n")
	}
}
