Create a command-line task manager application that allows users to manage their tasks. Users should be able to add tasks, mark tasks as completed, list tasks, and remove tasks.

### 1. Project Structure
The project is structured as follows:

```
task-manager/
├── cmd/
│   └── task/
│       ├── add.go
│       ├── complete.go
│       ├── list.go
│       ├── remove.go
│       └── root.go
├── internal/
│   └── task/
│       ├── task.go
│       └── task_manager.go
├── data/
│   └── tasks.json
├── main.go
└── go.mod
```

- **`cmd/`**: Contains the command-line interface (CLI) commands.
- **`internal/`**: Contains the internal package for managing tasks.
- **`data/`**: Stores the `tasks.json` file where tasks are saved.
- **`main.go`**: The main entry point of the application.

### 2. `Task` Struct (in `internal/task/task.go`)
This file defines the `Task` struct, which represents a single task:

```go
package task

import "time"

type Task struct {
    ID          int
    Description string
    DueDate     time.Time
    Completed   bool
}
```

### 3. `TaskManager` (in `internal/task/task_manager.go`)
This file contains the `TaskManager` struct and methods for managing tasks:

- **`TaskManager`**: Manages a slice of `Task` structs.

#### Methods:
- **`LoadTasks(filename string) error`**: Loads tasks from a JSON file.
- **`SaveTasks(filename string) error`**: Saves tasks to a JSON file.
- **`AddTask(description string, dueDate time.Time)`**: Adds a new task.
- **`CompleteTask(taskID int)`**: Marks a task as completed.
- **`RemoveTask(taskID int)`**: Removes a task from the list.

### 4. Command-Line Interface (CLI) Commands (in `cmd/task/`)
The CLI is built using the Cobra library.

#### `root.go`
- Sets up the main command and adds subcommands.
- Initializes the `TaskManager`, loads tasks on startup, and saves tasks on exit.

#### Subcommands:
1. **`add.go`**: Adds a new task.
2. **`list.go`**: Lists all tasks.
3. **`complete.go`**: Marks a task as completed.
4. **`remove.go`**: Removes a task.

### 5. `main.go` (in the project root)
This file imports the `cmd/task` package and starts the CLI:

- Initializes the CLI application.

### How the Project Works:
1. **Adding Tasks**:
   - When you run `./task-manager add -d "Do the dishes" -t 2024-03-14`, it creates a new task with the provided description and due date.
   - The `AddTask` function in `task_manager.go` creates a new `Task` struct and appends it to the `TaskManager`'s list of tasks.
   - The `SaveTasks` function is then called to save the updated task list to `data/tasks.json`.

2. **Listing Tasks**:
   - Running `./task-manager list` lists all tasks.
   - The `ListTasks` function in `list.go` fetches the list of tasks from the `TaskManager` and prints them to the console.

3. **Completing Tasks**:
   - When you run `./task-manager complete 1`, it marks task with ID `1` as completed.
   - The `CompleteTask` function in `complete.go` sets the `Completed` field of the specified task to `true`.
   - The updated task list is then saved to `data/tasks.json`.

4. **Removing Tasks**:
   - Running `./task-manager remove 1` removes the task with ID `1` from the list.
   - The `RemoveTask` function in `remove.go` removes the task with the specified ID from the `TaskManager`'s list of tasks.
   - The updated task list is saved to `data/tasks.json`.

5. **Loading and Saving Tasks**:
   - The `LoadTasks` function in `task_manager.go` loads tasks from `data/tasks.json` when the program starts.
   - The `SaveTasks` function saves tasks to `data/tasks.json` when the program exits.

### Workflow:
1. Start the program with `./task-manager`.
2. Add tasks using `./task-manager add -d "Task description" -t 2024-03-14`.
3. List tasks with `./task-manager list`.
4. Complete a task with `./task-manager complete 1`.
5. Remove a task with `./task-manager remove 1`.
6. Tasks are saved to `data/tasks.json` and will persist between program runs.


``` Usage CMD
$ ./task-manager -h
Simple Task Manager CLI

Usage:
  task [command]

Available Commands:
  add         Add a new task
  complete    Mark a task as completed
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all tasks
  remove      Remove a task

Flags:
  -h, --help   help for task

Use "task [command] --help" for more information about a command.

```
