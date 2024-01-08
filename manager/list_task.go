package manager

import "fmt"

func ListTask(tasks *[]Task) {
	fmt.Println("Tasks in DB:")
	for index, task := range *tasks {
		fmt.Printf("%d. Task Name: %s, Task Status: %s\n", index+1, task.Name, task.Status)
	}
}
