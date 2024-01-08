package utils

import (
	"log"

	"github.com/samiptimalsena/gutudu/manager"
)

func GetTaskInDB() *[]manager.Task {
	DB_NAME := "db.txt"
	tasks, err := ReadFromDB(DB_NAME)
	if err != nil {
		log.Println("[ERROR] Error while reading the file: ", err)
	}
	return &tasks
}
