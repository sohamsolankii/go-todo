package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	newTodo := Todo{
		Title:       title,
		CreatedAt:   time.Now(),
		Completed:   false,
		CompletedAt: nil,
	}

	*todos = append(*todos, newTodo)
}

func (todos *Todos) validateIdx(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index!")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := todos.validateIdx(index)

	if t != nil {
		return t
	}

	todo := *todos
	*todos = append(todo[:index], todo[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validateIdx(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	t[index].Completed = !isCompleted

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	} else {
		t[index].CompletedAt = nil
	}

	return nil
}

func (todos *Todos) update(index int, newTitle string) error {
	t := *todos

	if err := t.validateIdx(index); err != nil {
		return err
	}

	t[index].Title = newTitle
	return nil
}

func (todos *Todos) list() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for idx, todo := range *todos {
		completed := "X"
		completedAt := ""

		if todo.Completed {
			completed = "✔"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(idx), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
