package main

import (
	"encoding/json"
	"os"
)

func ReadFile(filename string) ([]Task, error) {
	var tasks []Task

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return tasks, nil
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func WriteFile(tasks []Task, filename string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
