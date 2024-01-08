package manager

import (
	"log"
)

func GetStatus(taskName *string, tasks *[]Task) string {
	for _, task := range *tasks {
		if task.Name == *taskName {
			return task.Status
		}
	}
	log.Printf("[DEBUG] %s task not found in DB", *taskName)
	return "NONE"
}

func ChangeStatus(taskName *string, tasks *[]Task) {
	for index, task := range *tasks {
		if task.Name == *taskName {
			(*tasks)[index].Status = "Completed"
			return
		}
	}
	log.Printf("[DEBUG] %s task not found in DB", *taskName)
}
