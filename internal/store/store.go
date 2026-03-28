package store

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/catintrenches1234/tasks/internal/model"
)

func ensureStoreDir(path string) error {
	dir := filepath.Dir(path)
	return os.MkdirAll(dir, 0o700)
}

func AddTask(filepath string, description string) error {
	if err := ensureStoreDir(filepath); err != nil {
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := LockFile(f); err != nil {
		return err
	}
	defer UnlockFile(f)

	tasks, err := LoadTasks(filepath)
	if err != nil {
		return err
	}

	maxId := 0
	for _, t := range tasks {
		if t.ID > maxId {
			maxId = t.ID
		}
	}

	task := model.Task{
		ID:          maxId + 1,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
	}

	tasks = append(tasks, task)

	return SaveTasks(filepath, tasks)
}

func ListTasks(filepath string, showAll bool) ([]model.Task, error) {
	tasks, err := LoadTasks(filepath)
	if err != nil {
		return nil, err
	}

	if showAll {
		return tasks, nil
	}

	var filtered []model.Task
	for _, t := range tasks {
		if !t.Completed {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil
}

func CompleteTask(filepath string, id int) error {
	if err := ensureStoreDir(filepath); err != nil {
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = LockFile(f); err != nil {
		return err
	}
	defer UnlockFile(f)

	tasks, err := LoadTasks(filepath)
	if err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task not found")
	}

	return SaveTasks(filepath, tasks)
}

func DeleteTask(filepath string, id int) error {
	if err := ensureStoreDir(filepath); err != nil {
		return err
	}

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = LockFile(f); err != nil {
		return err
	}
	defer UnlockFile(f)

	tasks, err := LoadTasks(filepath)
	if err != nil {
		return err
	}

	var updated []model.Task
	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		updated = append(updated, t)
	}

	if !found {
		return fmt.Errorf("task not found")
	}

	return SaveTasks(filepath, updated)
}
