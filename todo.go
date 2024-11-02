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

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) Delete(index int) error {
	todo := *todos
	if err := todo.validateIndex(index); err != nil {
		return err
	}

	*todos = append(todo[:index], todo[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	todo := *todos
	if err := todo.validateIndex(index); err != nil {
		return err
	}

	isCompleted := todo[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		todo[index].CompletedAt = &completionTime
		// added by me
		todo[index].Completed = true
		// and below else also added by me
	} else {
		todo[index].Completed = false
		todo[index].CompletedAt = nil
	}

	// todo[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	todo := *todos
	if err := todo.validateIndex(index); err != nil {
		return err
	}

	todo[index].Title = title
	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")
	for index, todo := range *todos {
		completed := "❌"
		completedAt := ""
		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
