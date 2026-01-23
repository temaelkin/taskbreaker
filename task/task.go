package task

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func NewTask(name string) Task {
	return Task{
		Name: name,
		Done: false,
	}
}

func (t *Task) MarkDone() {
	t.Done = true
}
