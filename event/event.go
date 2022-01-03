package event

const (
	EventSetPanel EventId = 0

	EventLogin       EventId = 1
	EventLoginResult EventId = 2

	EventLogout EventId = 3

	EventAddTeacher       EventId = 4
	EventAddTeacherResult EventId = 5

	EventLoadTimeTable       EventId = 6
	EventLoadTimeTableResult EventId = 7

	EventQuerryTaecher       EventId = 8
	EventQuerryTaecherResult EventId = 9

	EventUpdateTeacherList EventId = 10

	EventLoading EventId = 11

	EventUpdate EventId = 12

	eventMax EventId = 13
)

type EventId int
type ReciverId int

type event struct {
	id       EventId
	receiver []func(data interface{})
}

var events [eventMax]event

func Go(id EventId, data interface{}) {
	for _, r := range events[id].receiver {
		if r != nil {
			r(data)
		}
	}
}

func On(id EventId, f func(data interface{})) ReciverId {
	for i, r := range events[id].receiver {
		if r == nil {
			events[id].receiver[i] = f
			return (ReciverId)(i)
		}
	}

	events[id].receiver = append(events[id].receiver, f)
	return (ReciverId)(len(events[id].receiver) - 1)
}

func UnOn(id EventId, rId ReciverId) {
	if (ReciverId)(len(events[id].receiver)) <= rId {
		return
	}
	events[id].receiver[rId] = nil
}
