package concurrent

import "fmt"

type Agent struct {
	data        interface{}
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

func (a *Agent) Send(futureEdit func(interface{}) interface{})  {
	a.editQueue <- futureEdit
}

func (a *Agent) Get() interface{}  {
	return a.data
}

func (a *Agent) Alter(futureEdit func(interface{}) interface{}) *Future  {
	c := make(chan interface{}, 1)
	a.editQueue <- func (collection interface{}) interface{} {
		var editedCollection = futureEdit(collection)

		c <- editedCollection
		return editedCollection
	}

	return NewFuture(c)
}

// Get the current Collection of String after all the current edits have completed
func (a *Agent) Future() *Future {
	c := make(chan interface{}, 1)
	a.editQueue <- func(object interface{}) interface{} {
		c <- object
		return object
	}

	return NewFuture(c)
}

type Future struct {
	WaitChan chan interface{}
}

func NewFuture(c chan interface{}) *Future {
	future := &Future{c}
	return future
}




type FutureStringCollection struct {
	WaitChan chan []string
}

func NewFutureStringCollection(c chan []string) *FutureStringCollection {
	futureStringCollection := &FutureStringCollection{c}
	return futureStringCollection
}

type AgentStringCollection struct {
	stringCollection []string
	editQueue        chan func([]string) []string
}

func NewAgentStringCollection() *AgentStringCollection {
	agent := &AgentStringCollection{nil,
		make(chan func([]string) []string, 10)}
	go func() {
		for f := range agent.editQueue {
			agent.stringCollection = f(agent.stringCollection)
			fmt.Printf("update: %v\n", agent.stringCollection)
		}
	}()
	return agent
}

// Get the current Collection of String after all the current edits have completed
func (a *AgentStringCollection) Future() *FutureStringCollection {
	c := make(chan []string, 1)
	a.editQueue <- func(stringCollection []string) []string {
		c <- stringCollection
		return stringCollection
	}

	return NewFutureStringCollection(c)
}

func (a *AgentStringCollection) Send(futureEditCollection func([]string) []string)  {
	a.editQueue <- futureEditCollection
}

func (a *AgentStringCollection) Alter(futureEditCollection func([]string) []string) *FutureStringCollection  {
	c := make(chan []string, 1)
	a.editQueue <- func (collection []string) []string {
		var editedCollection = futureEditCollection(collection)

		c <- editedCollection
		return editedCollection
	}

	return NewFutureStringCollection(c)
}

func (a *AgentStringCollection) ContainsString(newString string) bool {
	return ContainsString(a.stringCollection, newString)
}

func (a *AgentStringCollection) StringCollection() []string {
	return a.stringCollection
}

func ContainsString(stringCollection []string, testString string) bool {
	for _, stringInCollection := range stringCollection {
		if stringInCollection == testString {
			return true
		}
	}

	return false
}

func (a *AgentStringCollection) AddUrlAsync(newString string) {
	a.editQueue <- func(stringCollection []string) []string {
		for _, url := range stringCollection {
			if url == newString {
				return stringCollection
			}
		}

		return append(stringCollection, newString)
	}
}
