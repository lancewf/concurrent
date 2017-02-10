package concurrent

import (
	"testing"
	"time"
)


func TestActorWorker(t *testing.T){
	actorWorker := NewActorWorker()

	replyChan := make(chan interface{}, 1)

	responseCount := 0
	responses := []string{"test", "testbob", "testbobtimmie", "testbobtimmiecraig"}

	go func() {
		for reply := range replyChan {
			switch v := reply.(type){
			case DataResponse:
				if v.data != responses[responseCount] {
					t.Error("Expected ", responses[responseCount], "got", v)
				}

				responseCount++
			}

		}
	}()

	actorWorker.Send <- Request{DataRequest{}, replyChan}

	actorWorker.Send <- Request{AppendRequest{"bob"}, replyChan}

	actorWorker.Send <- Request{DataRequest{}, replyChan}

	actorWorker.Send <- Request{AppendRequest{"timmie"}, replyChan}

	actorWorker.Send <- Request{DataRequest{}, replyChan}

	actorWorker.Send <- Request{AppendRequest{"craig"}, replyChan}

	actorWorker.Send <- Request{DataRequest{}, replyChan}

	time.Sleep(200 * time.Millisecond)

	if responseCount != 4 {
		t.Error("Expected 3 responses, got ", responseCount)
	}

}


