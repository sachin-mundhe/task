package cmd

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/spf13/cobra"

	db "../db"
	task "../task"
)

// import ( task "../../task")

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds task to your task list",
	Run:   addTask,
}

func init() {
	RootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {

	taskStr := strings.Join(args, " ")
	genID := bson.NewObjectId()
	newTask := task.Task{Todo: taskStr, ID: genID}

	db.Connect()

	err := db.TaskCollection.Insert(newTask)
	if err != nil {
		fmt.Println("Unable to save to DB")
	} else {
		fmt.Println("Task added successfully with ID:", genID.Hex())
	}

}
