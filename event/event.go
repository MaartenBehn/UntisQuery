package event

const (
	EventUpdate EventId = 0
	EventDraw   EventId = 1

	EventSetPage EventId = 2

	EventLogin       EventId = 3
	EventLoginResult EventId = 4
	EventLogout      EventId = 5

	EventAddTeacher        EventId = 6
	EventUpdateQuerryPanel EventId = 7

	EventLoadTimeTable EventId = 8

	EventQuerryTaecher EventId = 9

	EventStartLoading  EventId = 10
	EventUpdateLoading EventId = 11

	eventMax = 12
)

type EventId int
type ReciverId int

type event struct {
	id       EventId
	receiver []func(data interface{})
}

var events [eventMax]event

func Init() {
	for _, e := range events {
		e.receiver = []func(interface{}){}
	}
}

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
