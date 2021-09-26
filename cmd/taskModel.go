package cmd

type Task struct {
	Done bool   `json:"done,"`
	Desc string `json:"desc"`
}
