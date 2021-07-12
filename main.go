package main

import (
	Untis "UntisAPI"
	"fmt"
	"time"
)

var user *Untis.User
var timetable []map[int]Untis.Period
var rooms map[int]Untis.Room
var classes map[int]Untis.Class

func main() {
	user = Untis.NewUser("maarten8", "behn500", "TBZ Mitte Bremen", "https://tipo.webuntis.com")
	user.Login()

	date := Untis.ToUntisDate(time.Now())
	loadTimetable(date, date)

	querryTeacher("Daniel", "Dibbern")

	user.Logout()
}

func loadTimetable(startDate int, endDate int) {
	classes = user.GetClasses()
	rooms = user.GetRooms()

	timetable = []map[int]Untis.Period{}
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

	var foundPeriod []Untis.Period
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

		fmt.Print("\nClass: ")
		for _, classId := range period.Classes {
			fmt.Printf("%s ", classes[classId].Name)
		}

		date := Untis.ToGoDate(period.Date)
		fromTime := Untis.ToGoTime(period.StartTime)
		tillTime := Untis.ToGoTime(period.EndTime)

		fmt.Printf("\nDate: %d %s %d From: %02d:%02d  Till: %02d:%02d ",
			date.Day(), date.Month(), date.Year(),
			fromTime.Hour(), fromTime.Minute(),
			tillTime.Hour(), tillTime.Minute())

		fmt.Print("\n\n")
	}
}
