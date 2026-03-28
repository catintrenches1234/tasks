package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/catintrenches1234/tasks/internal/store"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Long: `Delete a task using its ID.

Examples:

	tasks delete 1
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID")
			os.Exit(1)
		}

		err = store.DeleteTask("~/.tasks/tasks.csv", id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Task deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
