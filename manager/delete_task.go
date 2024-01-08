package manager

import (
	"log"
)

func DeleteTask(taskName *string, tasks *[]Task) {
	tasksSlice := *tasks
	for index, task := range *tasks {
		if task.Name == *taskName {
			*tasks = append(tasksSlice[:index], tasksSlice[index+1:]...)
			return
		}
	}
	log.Printf("[DEBUG] %s task not found in DB", *taskName)
}
