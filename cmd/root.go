package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd RootCmd is starting point of the application
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
