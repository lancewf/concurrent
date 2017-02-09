package concurrent

import (
	"fmt"
)

type Vector struct {
	X int
	Y int
}

type Actor struct {
	Ask chan interface{}
}

func NewActor() *Actor {
	actor := &Actor{make(chan interface{}, 5)}
	actor.receive()
	return actor
}

func (a *Actor) receive() {
	go func() {
		for i := range a.Ask {
			switch v := i.(type) {
			case int:
				fmt.Printf("Twice %v is %v\n", v, v*2)
			case string:
				fmt.Printf("%q is %v bytes long\n", v, len(v))
			case Vector:
				fmt.Printf("x: %v, y:%v \n", v.X, v.Y)
			default:
				fmt.Printf("I don't know about type %T!\n", v)
			}
		}
	}()
}