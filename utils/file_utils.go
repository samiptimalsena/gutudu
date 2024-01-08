package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/samiptimalsena/gutudu/manager"
)

func WriteToDB(filename string, task []manager.Task) error {
	obj_path := filepath.Join("objs", filename)
	file, err := os.Create(obj_path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(task)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromDB(filename string) ([]manager.Task, error) {
	var task []manager.Task
	obj_path := filepath.Join("objs", filename)

	file, err := os.Open(obj_path)
	if err != nil {
		return task, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&task)
	if err != nil {
		return task, err
	}

	return task, nil
}
