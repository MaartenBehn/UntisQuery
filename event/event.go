package event

const (
	EventUpdate EventId = 1
	EventDraw   EventId = 2

	EventSetPage EventId = 3

	eventMax = 255
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
