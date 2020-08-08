package cmd

import (
	"fmt"
	"os"
	"task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks() //get the slice od tasks
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		//If no tasks are pending
		if len(tasks) == 0 {
			fmt.Println("No tasks pending")
			return
		}

		fmt.Println("Your TODO list: ")
		for i, task := range tasks {
			fmt.Printf("%v. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
