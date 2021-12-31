package untis

import (
	"fmt"
	"github.com/Stroby241/UntisAPI"
	"github.com/Stroby241/UntisQuerry/event"
	"github.com/Stroby241/UntisQuerry/state"
)

func addTeacher(firstname string, lastname string) bool {
	found := false
	for _, t2 := range state.Teachers {
		if t2.Firstname == firstname && t2.Lastname == lastname {
			found = true
		}
	}
	if found {
		return true
	}

	id, err := user.GetPersonId(firstname, lastname, true)
	if err != nil {
		return false
	}

	t := &state.Teacher{
		Firstname: firstname,
		Lastname:  lastname,
		Id:        id,
	}
	state.Teachers = append(state.Teachers, t)
	return true
}

func queryTeacher(teacher *state.Teacher) bool {
	if teacher == nil || timetable == nil || rooms == nil || classes == nil {
		return false
	}

	event.Go(event.EventStartLoading, "Scanning Data")
	state.FoundPeriods = []UntisAPI.Period{}
	for i, periods := range timetable {
		event.Go(event.EventUpdateLoading, float64(i)/float64(len(timetable))*100.0)
		for _, period := range periods {
			for _, testTeacher := range period.Teacher {
				if testTeacher == teacher.Id {
					state.FoundPeriods = append(state.FoundPeriods, period)
				}
			}
		}
	}

	fmt.Printf("\n%s %s found in %d Periods.\n", teacher.Firstname, teacher.Lastname, len(state.FoundPeriods))
	for _, period := range state.FoundPeriods {
		fmt.Printf("Room: ")
		for _, roomId := range period.Rooms {
			fmt.Printf("%s ", rooms[roomId].Name)
		}

		fmt.Printf("Class: ")
		for _, classId := range period.Classes {
			fmt.Printf("%s ", classes[classId].Name)
		}

		date := UntisAPI.ToGoDate(period.Date)
		fromTime := UntisAPI.ToGoTime(period.StartTime)
		tillTime := UntisAPI.ToGoTime(period.EndTime)

		fmt.Printf("Date: %d %s %d From: %02d:%02d  Till: %02d:%02d ",
			date.Day(), date.Month(), date.Year(),
			fromTime.Hour(), fromTime.Minute(),
			tillTime.Hour(), tillTime.Minute())

		fmt.Printf("\n")
	}

	event.Go(event.EventSetPage, state.PageQuerry)
	return true
}
