package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/catintrenches1234/tasks/internal/store"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Long: `Add creates a new task and stores it in the underlying data file.

The command takes a single argument: the task description. A unique ID is
automatically assigned, along with the current timestamp. The task is marked
as not completed by default.

Examples:

	tasks add "Tidy my desk"
	tasks add "Write project report"
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")

		err := store.AddTask("tasks.csv", description)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Task added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
