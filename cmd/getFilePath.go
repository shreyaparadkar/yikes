package cmd

import (
	"os"
	"path/filepath"
)

func getListFilePath() (string, error) {
	newpath, err := getDirPath()
	if err != nil {
		return "", err
	}
	file_path := filepath.Join(newpath, "todos.json")
	return file_path, nil
}

func getDirPath() (string, error) {
	home, err := os.UserHomeDir()
	dirpath := filepath.Join(home, "yikes")
	err = os.MkdirAll(dirpath, os.ModePerm)

	if err != nil {
		return "", err
	}
	return dirpath, nil
}
