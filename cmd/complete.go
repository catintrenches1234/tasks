package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/catintrenches1234/tasks/internal/store"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark a task as completed",
	Long: `Mark a task as completed using its ID.

Examples:

	tasks complete 1
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid task ID")
			os.Exit(1)
		}

		err = store.CompleteTask("~/.tasks/tasks.csv", id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println("Task marked as completed")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
