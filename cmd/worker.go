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

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start a worker registering all available workflows",
	Long: `This will start a worker registering all available workflows.
The workflows which will be registered can be modified from the code.
For getting details about workflows, read the docs in specific workflow folders.`,
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Perform start action",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		localworker.StartWorker(path)
	},
}

func init() {
	usersHome, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintln("Error occurred while getting user's home directory:", err))
	}
	usersHome = filepath.Join(usersHome)
	rootCmd.AddCommand(workerCmd)
	workerCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("path", "p", usersHome, "Folder location where config file is present")
}
