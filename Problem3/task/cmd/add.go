package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		argList := strings.Join(args, " ")
		fmt.Printf("Task added: %s\n", argList)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
