package cmd

import "github.com/spf13/cobra"

//RootCmd is the code to show the help list while running task in the CLI
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is your rudimentary CLI TODO manager",
}
