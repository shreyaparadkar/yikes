package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display list of all the tasks",
	Long: `Displays a list of all the tasks.
Filter the list based on the completion of tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		sortI, _ := cmd.Flags().GetBool("incomplete")
		sortC, _ := cmd.Flags().GetBool("complete")
		param := func() string {
			if sortI {
				return "i"
			} else if sortC {
				return "c"
			} else {
				return " "
			}
		}()
		showList(param)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("incomplete", "i", false, "Get only incomplete tasks")
	listCmd.Flags().BoolP("complete", "c", false, "Get only completed tasks")
}

func showList(sort string) error {
	const (
		checkbox = "☐"
		checked  = "✅"
	)
	file_path, err := getListFilePath()
	var tasks []Task
	stream, err := ioutil.ReadFile(file_path)
	if err != nil || len(stream) == 0 {
		fmt.Println("No tasks added yet!")
	}
	err = json.Unmarshal(stream, &tasks)
	if sort == "i" {
		for i, v := range tasks {
			if !v.Done {
				fmt.Print(i+1, ": ")
				fmt.Print(checkbox)
				fmt.Println(" ", v.Desc)
			}
		}
	} else if sort == "c" {
		for i, v := range tasks {
			if v.Done {
				fmt.Print(i+1, ": ")
				fmt.Print(checked)
				fmt.Println(" ", v.Desc)
			}
		}
	} else {
		for i, v := range tasks {
			fmt.Print(i+1, ": ")
			if v.Done {
				fmt.Print(checked)
			} else {
				fmt.Print(checkbox)
			}
			fmt.Println(" ", v.Desc)
		}
	}
	return nil
}
