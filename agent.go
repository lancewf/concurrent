package concurrent

type Agent struct {
	data      interface{}
	editQueue chan func(interface{}) interface{}
}

func NewAgent(initialData interface{}) *Agent {
	agent := &Agent{initialData,
		make(chan func(interface{}) interface{}, 5)}
	go func() {
		for f := range agent.editQueue {
			agent.data = f(agent.data)
		}
	}()
	return agent
}

func (a *Agent) Send(futureEdit func(interface{}) interface{}) {
	a.editQueue <- futureEdit
}

func (a *Agent) Get() interface{} {
	return a.data
}

func (a *Agent) Alter(futureEdit func(interface{}) interface{}) chan interface{} {
	c := make(chan interface{}, 1)
	a.editQueue <- func(collection interface{}) interface{} {
		var editedCollection = futureEdit(collection)

		c <- editedCollection
		return editedCollection
	}

	return c
}

// Get the current Collection of String after all the current edits have completed
func (a *Agent) Future() chan interface{} {
	c := make(chan interface{}, 1)
	a.editQueue <- func(object interface{}) interface{} {
		c <- object
		return object
	}

	return c
}


