package localworker

import "github.com/kasattejaswi/uberCadence-project/workflows/helloworld"

type Registries struct {
	registry interface{}
	alias    string
}

func GetWorkflowsForRegistration() []Registries {
	rlist := make([]Registries, 0, 10)
	rlist = append(rlist, Registries{
		registry: helloworld.HelloWorldWorkflow,
		alias:    helloworld.HelloWorldWorkflowName,
	})
	return rlist
}

func GetActivitiesForRegistration() []Registries {
	rlist := make([]Registries, 0, 10)
	rlist = append(rlist, Registries{
		registry: helloworld.HelloWorldActivity,
		alias:    "",
	})
	return rlist
}
