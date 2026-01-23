package storage

import (
	"encoding/json"
	"os"

	"github.com/temaelkin/taskbreaker/task"
)

type JSONStrorage struct {
	filename string
}

func NewJSON(filename string) *JSONStrorage {
	return &JSONStrorage{filename: filename}
}

func (s *JSONStrorage) Load() ([]task.Task, error) {
	file, err := os.Open(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []task.Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *JSONStrorage) Save(tasks []task.Task) error {
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	return encoder.Encode(tasks)
}
