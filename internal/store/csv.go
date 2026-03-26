package store

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/catintrenches1234/tasks/internal/model"
)

func LoadTasks(filepath string) ([]model.Task, error) {
	f,err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	var tasks []model.Task

	if _, err := r.Read(); err != nil {
		if err == io.EOF {
			return tasks, err
		}
		return nil, err
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) < 4 {
			return nil, fmt.Errorf("invalid record length")
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		description := record[1]

		createdAt, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			return nil, err
		}

		completed, err := strconv.ParseBool(record[3])
		if err != nil {
			return nil, err
		}

		task := model.Task{
			ID: id,
			Description: description,
			CreatedAt: createdAt,
			Completed: completed,
		}

		tasks = append(tasks, task)
	}

	return  tasks, nil
}

func SaveTasks(filepath string, tasks []model.Task) error {

	return nil
}
