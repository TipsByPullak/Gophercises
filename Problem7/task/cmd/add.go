package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		//Extract the task from the CLI args
		argList := strings.Join(args, " ")

		//Create the task and add it to the DB
		_, err := db.CreateTask(argList)
		if err != nil {
			fmt.Println("Error while adding task:", err.Error())
			os.Exit(1)
		}

		//Confirmation prompt
		fmt.Printf("Task added: %s\n", argList)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
