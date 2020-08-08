package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse", arg)
			}
			ids = append(ids, id) //Get all the ids to be marked done
		}
		tasks, err := db.AllTasks() //Fetch a list of all tasks
		if err != nil {
			fmt.Println("Something went wrong while removing:", err)
			os.Exit(1)
		}

		for _, id := range ids { //iterate over the tasks
			if id <= 0 || id > len(tasks) { //If invalid imput provided
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]             //Remember, list printed is 1-based, code is 0-based
			dErr := db.DeleteTask(task.Key) //Delete the task from the DB
			if dErr != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed\n", id)
			}
		}
		// fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
