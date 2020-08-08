package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir() //to initialise the DB on the homedir
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

//An error handler
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
