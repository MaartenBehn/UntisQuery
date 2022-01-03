package untis

import (
	"errors"
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

func queryTeacher(teacher *state.Teacher) {
	if teacher == nil || currentTimetable == nil || rooms == nil || classes == nil {
		event.Go(event.EventQuerryTaecherResult, errors.New("needed field are nil in queryTeacher"))
		return
	}

	event.Go(event.EventLoading, 0.0)
	foundPeriods := []UntisAPI.Period{}
	for i, period := range currentTimetable {
		event.Go(event.EventLoading, float64(i)/float64(len(timetable)))

		for _, testTeacher := range period.Teacher {
			if testTeacher == teacher.Id {
				foundPeriods = append(foundPeriods, period)
			}
		}
	}

	state.Periods = []*state.Period{}
	for i, foundPeriod := range foundPeriods {
		event.Go(event.EventLoading, float64(i)/float64(len(foundPeriods)))

		startTime := UntisAPI.ToGoTime(foundPeriod.StartTime)
		endTime := UntisAPI.ToGoTime(foundPeriod.EndTime)

		periodRooms := make([]string, len(foundPeriod.Rooms))
		for j, id := range foundPeriod.Rooms {
			periodRooms[j] = rooms[id].Name
		}
		periodClasses := make([]string, len(foundPeriod.Classes))
		for j, id := range foundPeriod.Classes {
			periodClasses[j] = classes[id].Name
		}
		periodSubjects := make([]string, len(foundPeriod.Subject))
		for j, id := range foundPeriod.Subject {
			periodSubjects[j] = subjects[id].Name
		}

		found := false
		for _, p := range state.Periods {

			theSame := false
			if len(p.Classes) == len(foundPeriod.Classes) &&
				len(p.Rooms) == len(foundPeriod.Rooms) &&
				len(p.Subjects) == len(foundPeriod.Subject) {
				theSame = true

				for _, class := range p.Classes {
					found := false
					for _, test := range periodClasses {
						if class == test {
							found = true
							continue
						}
					}

					if !found {
						theSame = false
					}
				}

				for _, room := range p.Rooms {
					found := false
					for _, test := range periodRooms {
						if room == test {
							found = true
							continue
						}
					}

					if !found {
						theSame = false
					}
				}

				for _, subject := range p.Subjects {
					found := false
					for _, test := range periodSubjects {
						if subject == test {
							found = true
							continue
						}
					}

					if !found {
						theSame = false
					}
				}
			}

			if theSame {
				if p.EndTime.Hour() == startTime.Hour() && p.EndTime.Minute() == startTime.Minute() {
					p.EndTime = endTime
					found = true
				} else if p.StartTime.Hour() == endTime.Hour() && p.StartTime.Minute() == endTime.Minute() {
					p.StartTime = startTime
					found = true
				}
			}
		}

		if !found {
			state.Periods = append(state.Periods, &state.Period{
				StartTime: startTime,
				EndTime:   endTime,
				Rooms:     periodRooms,
				Classes:   periodClasses,
				Subjects:  periodSubjects,
			})
		}
	}

	sort.Slice(state.Periods, func(i, j int) bool {
		return state.Periods[i].StartTime.Unix() < state.Periods[j].StartTime.Unix()
	})

	if len(state.Periods) == 0 {
		event.Go(event.EventQuerryTaecherResult, errors.New("no lessons for "+teacher.Firstname+" "+teacher.Lastname))
	}

	event.Go(event.EventLoading, 1.0)
	event.Go(event.EventQuerryTaecherResult, nil)
	return
}
