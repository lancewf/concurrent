package concurrent

type ActorReceiver interface {
	receive(request interface{}, sender chan<- interface{})
}

type Request struct {
	Data     interface{}
	Sender chan<- interface{}
}

type Actor struct {
	Send     chan<- Request
	receiver ActorReceiver
}

func NewActor(receiver ActorReceiver) *Actor {
	c := make(chan Request, 50)
	actor := &Actor{c, receiver}
	go func() {
		for request := range c {
			receiver.receive(request.Data, request.Sender)
		}
	}()
	return actor
}
