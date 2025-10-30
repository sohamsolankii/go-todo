# Go Todo CLI

A simple command line interface (CLI) application for managing your todos.

## Getting Started

To build the application, run the following command:

```bash
go build
```

## Usage

The following commands are available to manage your todos:

| Command      | Description                                            | Example                               |
|--------------|--------------------------------------------------------|---------------------------------------|
| `-add`       | Add a new todo.                                        | `go run ./ -add "Buy milk"`                |
| `-list`      | List all todos.                                        | `go run ./ -list`                          |
| `-edit`      | Edit an existing todo by its index.                    | `go run ./ -edit "1:Buy groceries"`        |
| `-toggle`    | Toggle the completion status of a todo by its index.   | `go run ./ -toggle 1`                      |
| `-del`       | Delete a todo by its index.                            | `go run ./ -del 1`                         |
