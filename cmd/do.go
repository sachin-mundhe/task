package cmd

import (
	"fmt"

	"gopkg.in/mgo.v2"

	db "../db"
	"github.com/spf13/cobra"
	"gopkg.in/mgo.v2/bson"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		db.Connect()
		for _, arg := range args {
			fmt.Println("abc", arg, "def")
			if isValid := bson.IsObjectIdHex(arg); isValid {

				if err := db.TaskCollection.RemoveId(bson.ObjectIdHex(arg)); err != nil {
					if err == mgo.ErrNotFound {
						fmt.Println("ID:", arg, "Task not found... Please enter valid ID")
					}
				} else {
					fmt.Println("Task successfully completed! ID:", arg)
				}

			} else {
				fmt.Println("Invalid ID:", arg)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
