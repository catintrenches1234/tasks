package model

import "time"

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	Completed   bool
}
