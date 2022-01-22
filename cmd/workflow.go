/*
Copyright © 2022 TEJASWI KASAT <kasattejasvi@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/kasattejaswi/uberCadence-project/localworker"
	"github.com/spf13/cobra"
)

// workflowCmd represents the workflow command
var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Helps in running workflows",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available workflows",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---------------------------------")
		fmt.Println("Below are the available workflows")
		fmt.Println("---------------------------------")
		localworker.PrintAvailableWorkflows()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a workflow using name or run all available workflows",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running workflows")
	},
}

func init() {
	rootCmd.AddCommand(workflowCmd)
	workflowCmd.AddCommand(listCmd)
	workflowCmd.AddCommand(runCmd)
}
