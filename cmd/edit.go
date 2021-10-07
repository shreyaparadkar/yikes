package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit task description",
	Long: `Used to edit description of the tasks in the list.
Syntax is yikes edit [task no] [updated description]`,
	Run: func(cmd *cobra.Command, args []string) {
		editTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func editTodo(args []string) error {
	//check if the args are valid
	if len(args) == 0 {
		fmt.Println("Enter valid task")
		return nil
	}

	//get task no and updated description of task
	taskNo, _ := strconv.Atoi(args[0])
	newDesc := strings.Join(args[1:], " ")

	file_path, err := getListFilePath()

	var tasks []Task
	stream, _ := ioutil.ReadFile(file_path)
	err = json.Unmarshal(stream, &tasks)

	//check if valid task no
	//if yes, update the description
	if taskNo > 0 && taskNo <= len(tasks) {
		tasks[taskNo-1].Desc = newDesc
		fmt.Println("Updated task successfully!")
	} else {
		fmt.Println("Enter valid task number")
	}

	taskJson, _ := json.Marshal(tasks)
	err = ioutil.WriteFile(file_path, taskJson, 0644)

	if err != nil {
		fmt.Println("Error updating the task.Please try again!")
		return err
	}
	return nil
}
