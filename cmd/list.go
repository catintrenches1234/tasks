package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/catintrenches1234/tasks/internal/store"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long: `List displays tasks stored in the data file.

By default, only uncompleted tasks are shown. Use the --all flag
to display all tasks, including completed ones.

Examples:

	tasks list
    tasks list --all
`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := store.ListTasks("tasks.csv", showAll)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		if showAll {
			fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
		} else {
			fmt.Fprintln(w, "ID\tTask\tCreated")
		}

		for _, t := range tasks {
			created := timediff.TimeDiff(t.CreatedAt)

			if showAll {
				fmt.Fprintf(w, "%d\t%s\t%s\t%t\n", t.ID, t.Description, created, t.Completed)
			} else {
				fmt.Fprintf(w, "%d\t%s\t%s\n", t.ID, t.Description, created)
			}
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")
}
