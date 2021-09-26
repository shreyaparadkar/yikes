package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tasks to the list",
	Long: `Used to add tasks to the list.
Syntax is yikes add [description of the task you want to add]`,
	Run: func(cmd *cobra.Command, args []string) {
		addTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTodo(args []string) error {
	//check if the args are valid
	if len(args) == 0 {
		fmt.Println("Enter valid task")
		return nil
	}

	//create a todo item using the Task struct
	descText := strings.Join(args, " ")
	newTodo := Task{false, descText}

	file_path, err := getListFilePath()

	//create a list of Tasks
	//read from the todos.json and unmarshal the data
	//then append the newTodo and marshal the data, and write it back to the file
	var tasks []Task
	stream, _ := ioutil.ReadFile(file_path)
	err = json.Unmarshal(stream, &tasks)

	tasks = append(tasks, newTodo)
	taskJson, _ := json.Marshal(tasks)

	err = ioutil.WriteFile(file_path, taskJson, 0644)

	if err != nil {
		fmt.Println("Error adding task.Please try again!")
		return err
	}

	fmt.Println("Task added:", descText)
	return nil
}
