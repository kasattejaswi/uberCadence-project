/*
Copyright © 2022 TEJASWI KASAT <kasattejasvi@gmail.com>

*/
package cmd

import (
	"github.com/kasattejaswi/uberCadence-project/configs"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Writes a default configuration file under user's home directory",
	Long: `This command will write a bare minimum configuration file under user's home directory.
You can change the configuration location by passing a path.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		force, _ := cmd.Flags().GetBool("force")
		configs.WriteConfigFile(path, force)
	},
}

func init() {
	initCmd.Flags().StringP("path", "p", "", "Path where config file will be written. Default will be user's home directory.")
	initCmd.Flags().BoolP("force", "f", false, "Use force to replace existing configuration and generate a new one.")
	rootCmd.AddCommand(initCmd)
}
