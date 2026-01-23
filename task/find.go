package task

import "strings"

func find(tasks []Task, query string) int {
	queryWords := strings.Fields(query)

	for i, t := range tasks {
		name := strings.ToLower(t.Name)

		match := true
		for _, word := range queryWords {
			if !strings.Contains(name, word) {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}

	return -1
}
