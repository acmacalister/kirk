package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	version = "0.0.1"
)

func main() {
	var kirkCmd = &cobra.Command{
		Use: "kirk",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the current version of kirk",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version", version)
		},
	}

	var deployCmd = &cobra.Command{
		Use:   "deploy [environment]",
		Short: "Deploy runlist for environment",
		Long:  "A bunch of nonsense about how to deploy stuff",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				config := loadConfig(args[0])
				runTaskList(&config)
			}
		},
	}

	kirkCmd.ParseFlags(os.Args)
	kirkCmd.AddCommand(versionCmd)
	kirkCmd.AddCommand(deployCmd)
	kirkCmd.Execute()
}
