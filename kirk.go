package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	version = "0.0.1"
)

var (
	Source string
)

func main() {
	var kirkCmd = &cobra.Command{
		Use: "kirk",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hi there")
			cmd.Help()
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of Kirk",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version", version)
		},
	}

	var deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "deploy's a spock file to server role list",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deploy args are:", args)
			cmd.Help()
		},
	}
	deployCmd.Flags().StringVarP(&Source, "source", "s", "stuff", "Source directory to read from")

	kirkCmd.ParseFlags(os.Args)
	fmt.Println("flags are:", kirkCmd.Root())
	kirkCmd.AddCommand(versionCmd)
	kirkCmd.AddCommand(deployCmd)
	kirkCmd.Execute()
}
