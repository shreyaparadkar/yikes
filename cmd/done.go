package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark the task as done",
	Long: `Used to mark the task as done
Specify the task number to mark it as done`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		param := func() bool {
			if all {
				return true
			} else {
				return false
			}
		}()
		markAsDone(args, param, true)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
	doneCmd.Flags().BoolP("all", "a", false, "Completed all tasks")
}

func markAsDone(args []string, all bool, done bool) error {
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
		fmt.Println("Error updating the task.Please try again!")
		return err
	}
	var taskNo int
	if !all {
		taskNo, _ = strconv.Atoi(args[0])
		taskNo -= 1
		for i, v := range tasks {
			if i == taskNo {
				temp := Task{done, v.Desc}
				updatedTasks = append(updatedTasks, temp)
			} else {
				updatedTasks = append(updatedTasks, v)
			}
		}
	} else {
		for _, v := range tasks {
			temp := Task{done, v.Desc}
			updatedTasks = append(updatedTasks, temp)
		}
	}

	updatedTaskJson, _ := json.Marshal(updatedTasks)

	err = ioutil.WriteFile(file_path, updatedTaskJson, 0644)

	if err != nil {
		fmt.Println("Error updating task.Please try again!")
		return err
	} else {
		fmt.Println("Task updated successfully!")
		showList(" ")
	}

	return nil
}
