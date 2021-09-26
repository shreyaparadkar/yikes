package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete tasks from a list",
	Long: `Delete a single or all tasks from a list
Specify the task number to delete it`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		param := func() bool {
			if all {
				return true
			} else {
				return false
			}
		}()
		deleteTask(args, param)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolP("all", "a", false, "Delete all the tasks")
}

func deleteTask(args []string, all bool) error {
	if len(args) == 0 && !all {
		fmt.Println("Enter valid task")
		return nil
	}
	file_path, err := getListFilePath()
	var tasks []Task
	var updatedTasks []Task
	stream, _ := ioutil.ReadFile(file_path)
	err = json.Unmarshal(stream, &tasks)

	if err != nil {
		fmt.Println("Error delete the task.Please try again!")
		return err
	}
	var taskNo int
	if !all {
		taskNo, _ = strconv.Atoi(args[0])
		taskNo -= 1
		for i, v := range tasks {
			if i == taskNo {
				continue
			} else {
				updatedTasks = append(updatedTasks, v)
			}
		}
	}
	// updatedTaskJson, _ := json.Marshal(updatedTasks)
	var updatedTaskJson []byte

	if all {
		updatedTaskJson = []byte("")
	} else {
		updatedTaskJson, _ = json.Marshal(updatedTasks)
	}

	err = ioutil.WriteFile(file_path, updatedTaskJson, 0644)

	if err != nil {
		fmt.Println("Error deleting task.Please try again!")
		return err
	} else {
		fmt.Println("Task deleted successfully!")
	}

	return nil
}
