package task

import "errors"

type Manager struct {
	tasks []Task
}

func NewManager(tasks []Task) *Manager {
	return &Manager{tasks: tasks}
}

func (m *Manager) Add(name string) {
	m.tasks = append(m.tasks, NewTask(name))
}

func (m *Manager) All() []Task {
	return m.tasks
}

func (m *Manager) Done(query string) error {
	index := find(m.tasks, query)
	if index == -1 {
		return errors.New("Not found")
	}
	m.tasks[index].MarkDone()
	return nil
}
