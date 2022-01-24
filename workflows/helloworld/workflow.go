package helloworld

import (
	"context"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

/**
 * This is the hello world workflow sample.
 */

// ApplicationName is the task list for this sample
const ApplicationName = "helloWorldGroup"

const HelloWorldWorkflowName = "helloWorldWorkflow"

// helloWorkflow workflow decider
// With this, you cant accept the parameters dynamically since it
// is generallised to every workflow. Would suggest to write a function
// which will provide parameters required
func HelloWorldWorkflow(ctx workflow.Context) error {
	name := HelloWorldParameters()
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("helloworld workflow started")
	var helloworldResult string
	err := workflow.ExecuteActivity(ctx, HelloWorldActivity, name).Get(ctx, &helloworldResult)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return err
	}
	logger.Info("Workflow completed.", zap.String("Result", helloworldResult))
	return nil
}

func HelloWorldActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("helloworld activity started")
	return "Hello " + name + "!", nil
}

func HelloWorldParameters() string {
	return "Tejaswi"
}
