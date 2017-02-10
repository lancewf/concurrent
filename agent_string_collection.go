package concurrent

type AgentStringCollection struct {
	agent *Agent
}

func NewAgentStringCollection(initialData []string) *AgentStringCollection {
	return &AgentStringCollection{NewAgent(initialData)}
}

func (asc *AgentStringCollection) GetStringCollection() []string {
	return asc.agent.Get().([]string)
}

func (asc *AgentStringCollection) FutureStringCollection() chan []string {
	c := make(chan []string, 1)

	go func(){
		c <- (<-asc.agent.Future()).([]string)
	}()

	return c
}

func (asc *AgentStringCollection) AlterStringCollection(futureEdit func([]string) []string) chan []string {
	c := make(chan []string, 1)

	var wrappedFunc = func(object interface{}) interface{} {
		return futureEdit(object.([]string))
	}

	go func(){
		c <- (<-asc.agent.Alter(wrappedFunc)).([]string)
	}()

	return c
}

func (asc *AgentStringCollection) SendStringCollection(futureEdit func([]string) []string) {

	var wrappedFunc = func(object interface{}) interface{} {
		return futureEdit(object.([]string))
	}

	asc.agent.Send(wrappedFunc)
}


