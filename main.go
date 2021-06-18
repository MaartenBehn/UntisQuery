package main

import (
	"UntisQuerry/UntisV2"
	"fmt"
)

var timetable []map[int]UntisV2.Period

func main() {
	user := UntisV2.NewUser("dummy2", "TBZ2020!x", "TBZ Mitte Bremen", "https://tipo.webuntis.com")
	user.Login()
	rooms := user.GetRooms()

	counter := 0
	for _, room := range rooms {
		fmt.Printf("Loading Room: %d of %d. \r", counter, len(rooms))

		periods := user.GetTimeTable(room.Id, 4, 20210512, 20210512)

		if periods != nil {
			timetable = append(timetable, periods)
		}

		counter++
	}
	id := user.GetPersonId("Daniel", "Dibbern")

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

	for _, period := range foundPeriod {

		for _, roomId := range period.Rooms {
			fmt.Printf("%s\n", rooms[roomId].Name)
		}
	}

	user.Logout()
}
