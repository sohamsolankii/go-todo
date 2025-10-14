package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlag struct {
	Add string
	Del int
	Edit string
	Toggle int
	List bool
}

func NewCmdFlag() *cmdFlag {
	cf := cmdFlag{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specifying a title.")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index.")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by index to toggle.")
	flag.BoolVar(&cf.List, "list", false, "List all todos.")

	flag.Parse()

	return &cf
}


func (cf *cmdFlag) execute(todos *Todos) {
	switch {
		case cf.List:
			todos.list()
		case cf.Add != "":
			todos.add(cf.Add)
		case cf.Edit != "":
			parts := strings.SplitN(cf.Edit, ":", 2)
			if len(parts) != 2 {
				fmt.Println("Invalid format for edit. Use id:title")
				os.Exit(1)
			}

			index, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Invalid id format for edit. Use id:title")
				os.Exit(1)
			}

			todos.update(index, parts[1])
		case cf.Del != -1:
			todos.delete(cf.Del)
		case cf.Toggle != -1:
			todos.toggle(cf.Toggle)
		default:
			fmt.Println("No valid command!")
			os.Exit(1)
	}
}