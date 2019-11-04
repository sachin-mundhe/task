package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/mgo.v2/bson"

	db "../db"
	task "../task"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Displays all tasks from your task list",
	Run: func(cmd *cobra.Command, args []string) {

		var allTasks []task.Task
		db.Connect()
		err := db.TaskCollection.Find(bson.M{}).All(&allTasks)
		if err != nil {
			fmt.Println("DB issue..")
			return
		}

		if len(allTasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		fmt.Println("Index \t ID \t\t\t\t Task")
		for i, v := range allTasks {
			fmt.Printf("%d \t %s \t %s\n", i, v.ID.Hex(), v.Todo)
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
