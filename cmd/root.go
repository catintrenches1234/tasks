package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Task manager CLI",
	Long: `tasks is a simple command-line application for managing tasks.

You can add, list, complete, and delete tasks stored in a local data file.

Examples:

	tasks add "Tidy my desk"
	tasks list
	tasks complete 1
	tasks delete 1
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

}
