/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start a worker registering all available workflows",
	Long: `This will start a worker registering all available workflows.

The workflows which will be registered can be modified from the code.

For getting details about workflows, read the docs in specific workflow folders.

Currently below workflows will be registered:
------------------------------------------------
helloworld`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("worker called")
	// },
}

func init() {
	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
