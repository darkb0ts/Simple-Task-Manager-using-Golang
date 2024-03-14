package task

import "time"

type Task struct {
	ID          int
	Description string
	DueDate     time.Time
	Completed   bool
}
