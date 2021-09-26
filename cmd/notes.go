package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Add notes",
	Long: `Add notes or other temp data that can be retrived later.
Add quotes("") around the title and desc flag when entering data`,
	Run: func(cmd *cobra.Command, args []string) {
		create, _ := cmd.Flags().GetBool("create")
		open, _ := cmd.Flags().GetBool("open")
		remove, _ := cmd.Flags().GetBool("remove")
		get, _ := cmd.Flags().GetBool("get")
		title, _ := cmd.Flags().GetString("title")
		desc, _ := cmd.Flags().GetString("desc")
		if create {
			createNote(title, desc, open)
		} else if !create && open {
			openNote(title)
		} else if remove {
			deleteNote(title)
		} else if get {
			getNotesList()
		} else {
			fmt.Println("Enter valid flags")
		}
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)
	notesCmd.Flags().BoolP("create", "c", false, "Create a note")
	notesCmd.Flags().BoolP("open", "o", false, "Open note")
	notesCmd.Flags().BoolP("remove", "r", false, "Delete note")
	notesCmd.Flags().BoolP("get", "g", false, "Get list of notes")
	notesCmd.Flags().StringP("title", "t", "", "Add the title of the txt file")
	notesCmd.Flags().StringP("desc", "d", "", "Add the contents of the txt file")
}

func getNotesPath(title string) (string, error) {
	dir_path, err := getDirPath()
	if err != nil {
		return "", err
	}
	fileName := "note-" + title + ".txt"
	file_path := filepath.Join(dir_path, fileName)
	return file_path, nil
}

func createNote(title string, desc string, open bool) error {
	if len(title) == 0 {
		fmt.Println("Enter valid title of the note")
		return nil
	}
	file_path, err := getNotesPath(title)
	file, err := os.Create(file_path)

	if err != nil {
		fmt.Println("Error creating the note.Please try again!")
		return err
	}
	file.WriteString(desc)
	defer file.Close()
	if open {
		openNote(title)
	} else {
		fmt.Println("New note created!")
	}
	return nil
}

func openNote(title string) error {
	if len(title) == 0 {
		fmt.Println("Enter valid filename")
	}
	file_path, err := getNotesPath(title)
	if err != nil {
		return err
	}
	cmd := exec.Command(`explorer`, `/open,`, file_path)
	cmd.Run()
	return nil
}

func deleteNote(title string) error {
	file_path, err := getNotesPath(title)
	err = os.Remove(file_path)
	if err != nil {
		fmt.Println("Error deleting the note")
		return err
	}
	fmt.Println("Note deleted!")
	return nil
}

func getNotesList() error {
	dir_path, err := getDirPath()
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(dir_path)
	if err != nil {
		fmt.Println("Error getting the notes")
		return err
	}
	c := 0
	for _, file := range files {
		temp := strings.Split(file.Name(), ".")
		temp1 := strings.Split(temp[0], "-")
		if len(temp) == 2 && temp[1] == "txt" && len(temp1) >= 2 && temp1[0] == "note" {
			fmt.Println(file.Name())
			c++
		}
	}
	if c == 0 {
		fmt.Println("No notes found!")
	}
	return nil
}
