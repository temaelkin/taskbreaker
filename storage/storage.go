package storage

import (
	"github.com/temaelkin/taskbreaker/task"
)

type TaskStorage interface {
	Save([]task.Task) error
	Load() ([]task.Task, error)
}
