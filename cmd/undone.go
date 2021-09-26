package cmd

import (
	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "Mark a task as not-done",
	Long: `Used to mark tasks as not-done
Specify the task number to mark it as undone`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		param := func() bool {
			if all {
				return true
			} else {
				return false
			}
		}()
		markAsDone(args, param, false)
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
	undoneCmd.Flags().BoolP("all", "a", false, "Completed all tasks")
}
