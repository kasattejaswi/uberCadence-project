package localworker

import (
	"fmt"
	"path/filepath"

	"github.com/kasattejaswi/uberCadence-project/helper"
	"github.com/kasattejaswi/uberCadence-project/statics"
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
	h.StartWorkers(h.Config.DomainName, "uberCadence-project", workerOptions)
}
