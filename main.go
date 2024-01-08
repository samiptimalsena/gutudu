package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/samiptimalsena/gutudu/manager"
	"github.com/samiptimalsena/gutudu/utils"
)

func main() {
	log.Println("[INFO] Welcome to GUTUDU")

	tasks := utils.GetTaskInDB()

	addTask := flag.String("add", "", "Name of task to add.")
	delTask := flag.String("del", "", "Name of task to remove.")
	statusTask := flag.String("status", "", "Name of task to check the status")
	changeStatus := flag.String("change", "", "Name of task to change the status")
	listTask := flag.Bool("list", false, "Bool to list all the tasks")
	flag.Parse()

	if *addTask != "" {
		manager.AddTask(addTask, tasks)
	} else if *delTask != "" {
		manager.DeleteTask(delTask, tasks)
	} else if *statusTask != "" {
		status := manager.GetStatus(statusTask, tasks)
		fmt.Printf("Task %s has %s status\n", *statusTask, status)
	} else if *changeStatus != "" {
		manager.ChangeStatus(changeStatus, tasks)
	} else if *listTask {
		manager.ListTask(tasks)
	}
	utils.WriteToDB("db.txt", *tasks)
}
