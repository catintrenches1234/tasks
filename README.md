# Todo CLI

A command line todo application for managing tasks in a local CSV data store.

## Overview

This project provides a terminal-based task manager with support for:

- adding new tasks
- listing uncompleted tasks
- listing all tasks with a `--all` flag
- marking tasks complete
- deleting tasks
- storing tasks in `~/.tasks/tasks.csv`
- file locking to prevent concurrent read/write access

## Commands

### Add a task

Create a new task with a description:

```bash
$ tasks add "My new task"
```

### List tasks

Show incomplete tasks:

```bash
$ tasks list
```

Show all tasks, including completed ones:

```bash
$ tasks list --all
```

### Complete a task

Mark a task as done by its ID:

```bash
$ tasks complete <taskid>
```

### Delete a task

Remove a task from the store:

```bash
$ tasks delete <taskid>
```

## Data storage

Tasks are stored in a CSV file at `~/.tasks/tasks.csv` with the following columns:

- `ID`
- `Description`
- `CreatedAt`
- `Completed`

The application resolves `~` to the current user home directory and creates the store directory automatically when needed.

## Packages used

- `encoding/csv` for CSV storage
- `strconv` for parsing and formatting numeric and boolean values
- `text/tabwriter` for aligned command output
- `os` for file and path operations
- `github.com/spf13/cobra` for the CLI interface
- `github.com/mergestat/timediff` for friendly relative timestamps

## Build and run

From the repository root:

```bash
go build -o tasks .
```

Then run the CLI:

```bash
./tasks <command>
```

## Notes

- the file store is locked during write operations to avoid concurrent access issues
