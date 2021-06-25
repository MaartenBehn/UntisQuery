package main

import (
	"UntisQuerry/Untis"
	"fmt"
)

var user *Untis.User
var timetable []map[int]Untis.Period
var rooms map[int]Untis.Room
var classes map[int]Untis.Class

func main3() {

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

func querryTeacher(firstname string, lastname string) string {
	result := ""

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

	result += fmt.Sprintf("%s %s found in %d Periods\n", firstname, lastname, len(foundPeriod))
	for _, period := range foundPeriod {
		result += fmt.Sprintf("Room: ")
		for _, roomId := range period.Rooms {
			result += fmt.Sprintf("%s ", rooms[roomId].Name)
		}

		result += fmt.Sprintf("Class: ")
		for _, classId := range period.Classes {
			result += fmt.Sprintf("%s ", classes[classId].Name)
		}

		date := Untis.ToGoDate(period.Date)
		fromTime := Untis.ToGoTime(period.StartTime)
		tillTime := Untis.ToGoTime(period.EndTime)

		result += fmt.Sprintf("Date: %d %s %d From: %02d:%02d  Till: %02d:%02d ",
			date.Day(), date.Month(), date.Year(),
			fromTime.Hour(), fromTime.Minute(),
			tillTime.Hour(), tillTime.Minute())

		result += fmt.Sprintf("\n")
	}
	return result
}
