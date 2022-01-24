package localworker

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/kasattejaswi/uberCadence-project/helper"
	"github.com/kasattejaswi/uberCadence-project/statics"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
)

func StartWorker(path string) {
	path = filepath.Join(path, statics.ConfigFileName)
	fmt.Println("Reading configuration at:", path)
	var h helper.Helper
	h.SetupServiceConfig(path)
	registerWorkflowAndActivity(&h)
	fmt.Println("Starting worker")
	launchWorkers(&h)
	select {}
}

func StartAllWorkflows(path string) {
	path = filepath.Join(path, statics.ConfigFileName)
	fmt.Println("Reading configuration at:", path)
	var h helper.Helper
	h.SetupServiceConfig(path)
	fmt.Println("Starting all workflows")
	allWorkflows := GetWorkflowsForRegistration()
	for _, a := range allWorkflows {
		name := a.alias
		fmt.Println("Starting workflow", name)
		workflowOptions := client.StartWorkflowOptions{
			ID:                              name + "_" + uuid.New(),
			TaskList:                        statics.TaskListName,
			ExecutionStartToCloseTimeout:    time.Minute,
			DecisionTaskStartToCloseTimeout: time.Minute,
		}
		h.StartWorkflow(workflowOptions, name)
	}
}

func StartWorkflow(path string, name string) {
	path = filepath.Join(path, statics.ConfigFileName)
	fmt.Println("Reading configuration at:", path)
	var h helper.Helper
	h.SetupServiceConfig(path)
	fmt.Println("Starting workflow", name)
	workflowOptions := client.StartWorkflowOptions{
		ID:                              name + "_" + uuid.New(),
		TaskList:                        statics.TaskListName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	h.StartWorkflow(workflowOptions, name)
}

func registerWorkflowAndActivity(
	h *helper.Helper,
) {
	allWorkflows := GetWorkflowsForRegistration()
	allActivities := GetActivitiesForRegistration()
	for _, val := range allWorkflows {
		h.RegisterWorkflowWithAlias(val.registry, val.alias)
	}
	for _, val := range allActivities {
		h.RegisterActivityWithAlias(val.registry, val.alias)
	}
}

func launchWorkers(h *helper.Helper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
		FeatureFlags: client.FeatureFlags{
			WorkflowExecutionAlreadyCompletedErrorEnabled: true,
		},
	}
	h.StartWorkers(h.Config.DomainName, statics.TaskListName, workerOptions)
}
