# CLI Task Manager

A command-line to-do list manager built in Go. This application allows users to add, list, mark, and delete tasks directly from the terminal.

## Features

- **Add new tasks**: Create to-do items and add them to your list.
- **List tasks**: View all tasks with their status (pending or completed).
- **Mark tasks as complete**: Keep track of tasks you've finished.
- **Delete tasks**: Remove tasks you no longer need.

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/ZnarKhalil/todos
   cd todos

   ```

2. **Install dependencies**:
   Go modules should handle dependencies. Just run:

   ```bash
   go mod tidy
   ```

## Usage

Once built, you can use the todo executable to manage tasks. Below are some basic commands:

    # Add a new task
    go run ./ -add "Buy groceries"

    # List all tasks
    go run ./ -list

    # Edit task
    go run ./ -edit "0:Buy Milk"

    # Mark a task as complete
    go run ./ -toggle 0

    # Delete a task
    go run ./ -delete 0

## Code Structure

- **main.go**: Initializes the application and handles command-line arguments.
- **command.go**: Manages the various commands (add, list, complete, delete).
- **storage.go**: Handles saving and loading tasks from a file or database.
- **todo.go**: Defines the Todo structure and any related functions.
