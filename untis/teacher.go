package untis

import (
	"fmt"
	"github.com/Stroby241/UntisAPI"
	"github.com/Stroby241/UntisQuerry/event"
	"github.com/Stroby241/UntisQuerry/state"
	"sort"
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

type period struct {
	startTime int
	endTime   int
	classes   []int
	subjects  []int
	rooms     []int
}

func queryTeacher(teacher *state.Teacher) bool {
	if teacher == nil || timetable == nil || rooms == nil || classes == nil {
		return false
	}

	event.Go(event.EventStartLoading, "Scanning Data")
	foundPeriods := []UntisAPI.Period{}
	for i, periods := range timetable {
		event.Go(event.EventUpdateLoading, float64(i)/float64(len(timetable))*100.0)
		for _, period := range periods {
			for _, testTeacher := range period.Teacher {
				if testTeacher == teacher.Id {
					foundPeriods = append(foundPeriods, period)
				}
			}
		}
	}

	periods := []*period{}
	for _, foundPeriod := range foundPeriods {

		found := false
		for _, p := range periods {
			if p.endTime == foundPeriod.StartTime {
				p.endTime = foundPeriod.EndTime
				found = true
			} else if p.startTime == foundPeriod.EndTime {
				p.startTime = foundPeriod.StartTime
				found = true
			}
		}

		if !found {
			periods = append(periods, &period{
				startTime: foundPeriod.StartTime,
				endTime:   foundPeriod.EndTime,
				classes:   foundPeriod.Classes,
				subjects:  foundPeriod.Subject,
				rooms:     foundPeriod.Rooms,
			})
		}
	}

	sort.Slice(periods, func(i, j int) bool {
		return periods[i].startTime < periods[j].startTime
	})

	result := fmt.Sprintf("%s %s found in %d Periods\n", teacher.Firstname, teacher.Lastname, len(foundPeriods))
	for _, period := range periods {
		fromTime := UntisAPI.ToGoTime(period.startTime)
		tillTime := UntisAPI.ToGoTime(period.endTime)

		result += fmt.Sprintf("From: %02d:%02d  Till: %02d:%02d ",
			fromTime.Hour(), fromTime.Minute(),
			tillTime.Hour(), tillTime.Minute())

		result += fmt.Sprintf("Room: ")
		for _, roomId := range period.rooms {
			result += fmt.Sprintf("%s ", rooms[roomId].Name)
		}

		result += fmt.Sprintf("Class: ")
		for _, classId := range period.classes {
			result += fmt.Sprintf("%s ", classes[classId].Name)
		}

		result += fmt.Sprintf("\n")
	}
	event.Go(event.EventUpdateQuerryText, result)

	event.Go(event.EventSetPage, state.PageQuerry)
	return true
}
