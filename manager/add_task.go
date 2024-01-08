package manager

func AddTask(task_name *string, tasks *[]Task) {
	task := Task{
		Name:   *task_name,
		Status: "Started",
	}
	*tasks = append(*tasks, task)
}
