package untis

import (
	"errors"
	"fmt"
	"github.com/Stroby241/UntisAPI"
	"github.com/Stroby241/UntisQuery/event"
	"github.com/Stroby241/UntisQuery/state"
	"sort"
)

func addTeacher(firstname string, lastname string) {
	event.Go(event.EventLoading, 0.0)

	if firstname == "" || lastname == "" {
		event.Go(event.EventAddTeacherResult, errors.New("Firstname and Lastname must be set."))
		return
	}

	event.Go(event.EventLoading, 0.1)
	found := false
	for _, t2 := range state.Teachers {
		if t2.Firstname == firstname && t2.Lastname == lastname {
			found = true
		}
	}
	if found {
		event.Go(event.EventAddTeacherResult, nil)
		return
	}

	event.Go(event.EventLoading, 0.5)
	id, err := user.GetPersonId(firstname, lastname, true)
	if err != nil {
		event.Go(event.EventAddTeacherResult, err)
		return
	}

	t := &state.Teacher{
		Firstname: firstname,
		Lastname:  lastname,
		Id:        id,
	}
	state.Teachers = append(state.Teachers, t)

	event.Go(event.EventLoading, 1.0)
	event.Go(event.EventAddTeacherResult, nil)
	return
}

type period struct {
	startTime int
	endTime   int
	classes   []int
	subjects  []int
	rooms     []int
}

func queryTeacher(teacher *state.Teacher) {
	if teacher == nil || timetable == nil || rooms == nil || classes == nil {
		event.Go(event.EventQuerryTaecherResult, errors.New("needed field are nil in queryTeacher"))
		return
	}

	event.Go(event.EventLoading, 0.0)
	foundPeriods := []UntisAPI.Period{}
	for i, periods := range timetable {
		event.Go(event.EventLoading, float64(i)/float64(len(timetable)))

		for _, period := range periods {
			for _, testTeacher := range period.Teacher {
				if testTeacher == teacher.Id {
					foundPeriods = append(foundPeriods, period)
				}
			}
		}
	}

	periods := []*period{}
	for i, foundPeriod := range foundPeriods {
		event.Go(event.EventLoading, float64(i)/float64(len(foundPeriods)))

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
	for i, period := range periods {
		event.Go(event.EventLoading, float64(i)/float64(len(periods)))

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

	event.Go(event.EventLoading, 1.0)
	event.Go(event.EventQuerryTaecherResult, result)
	return
}
