/*
Copyright © 2022 TEJASWI KASAT <kasattejasvi@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		allFlag, _ := cmd.Flags().GetBool("all")
		nameFlag, _ := cmd.Flags().GetString("name")
		pathFlag, _ := cmd.Flags().GetString("path")
		if allFlag {
			localworker.StartAllWorkflows(pathFlag)
		} else if nameFlag != "" {
			localworker.StartWorkflow(pathFlag, nameFlag)
		} else {
			fmt.Println("ERROR: No workflow name passed")
		}
	},
}

func init() {
	rootCmd.AddCommand(workflowCmd)
	workflowCmd.AddCommand(listCmd)
	workflowCmd.AddCommand(runCmd)
	usersHome, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintln("Error occurred while getting user's home directory:", err))
	}
	usersHome = filepath.Join(usersHome)
	runCmd.Flags().BoolP("all", "a", false, "Use this flag to start all the workflows at once")
	runCmd.Flags().StringP("name", "n", "", "Name of workflow to run")
	runCmd.Flags().StringP("path", "p", usersHome, "Folder location where config file is present")

}
