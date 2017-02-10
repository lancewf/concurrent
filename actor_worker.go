package concurrent


type AppendRequest struct {
	data string
}

type DataRequest struct {
}

type DataResponse struct {
	data string
}

type actorReceiverImp struct {
	localData string
}

func (a *actorReceiverImp) receive(request interface{}, sender chan<- interface{}) {
	switch r := request.(type) {
	case DataRequest:
		sender <- DataResponse{a.localData}
	case AppendRequest:
		a.localData = a.localData + r.data
	}
}


func NewActorWorker() *Actor {
	return NewActor(&actorReceiverImp{"test"})
}
