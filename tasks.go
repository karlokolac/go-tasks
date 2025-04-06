package main

import (
	"fmt"
	"strings"
	"time"
)

/*
	task-cli add 	<string>			DONE
	task-cli update <int> <string>		DONE
	task-cli delete <int>				DONE

	task-cli mark-in-progress 	<int>	DONE
	task-cli mark-done 			<int>

	task-cli list						DONE
	task-cli list done					DONE
	task-cli list todo					DONE
	task-cli list in-progress			DONE
*/

type status string

const (
	todo       status = "todo"
	done       status = "done"
	inProgress status = "in progress"
)

type Task struct {
	ID          int
	Description string
	Status      status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CheckArgsLength(args []string, n int) bool {
	if len(args) != n {
		return false
	}
	return true
}

func GetLastTaskID(filename string) (int, error) {
	tasks, err := ReadFile(filename)
	if err != nil {
		return -1, err
	}

	if len(tasks) == 0 {
		return 0, nil
	} else {
		return tasks[len(tasks)-1].ID, nil
	}
}

func AddTask(description, filename string) (bool, string) {
	task := Task{
		Description: description,
		Status:      todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := GetLastTaskID(filename)
	if err != nil {
		return false, err.Error()
	}
	if id >= 0 {
		task.ID = id + 1
	} else {
		return false, "Failed to assign task ID"
	}

	tasks, err := ReadFile(filename)
	if err != nil {
		return false, err.Error()
	}

	tasks = append(tasks, task)
	err = WriteFile(tasks, filename)
	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

func UpdateTask(id int, description, filename string) (bool, string) {
	tasks, err := ReadFile(filename)
	if err != nil {
		return false, err.Error()
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			updated = true
		}
	}

	if updated {
		err = WriteFile(tasks, filename)
		if err != nil {
			return false, err.Error()
		}
	} else {
		return false, "Task with that ID not found"
	}

	return true, ""
}

func DeleteTask(id int, filename string) (bool, string) {
	tasks, err := ReadFile(filename)
	if err != nil {
		return false, err.Error()
	}

	deleted := false
	var NotDeletedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			deleted = true
			continue
		} else {
			NotDeletedTasks = append(NotDeletedTasks, task)
		}
	}

	if deleted {
		err = WriteFile(NotDeletedTasks, filename)
		if err != nil {
			return false, err.Error()
		}
	} else {
		return false, "Task with that ID not found"
	}

	return true, ""
}

func MarkTaskInProgress(id int, filename string) (bool, string) {
	tasks, err := ReadFile(filename)
	if err != nil {
		return false, err.Error()
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "in progress"
			tasks[i].UpdatedAt = time.Now()
			updated = true
		}
	}

	if updated {
		err = WriteFile(tasks, filename)
		if err != nil {
			return false, err.Error()
		}
	} else {
		return false, "Task with that ID not found"
	}

	return true, ""
}

func MarkTaskDone(id int, filename string) (bool, string) {
	tasks, err := ReadFile(filename)
	if err != nil {
		return false, err.Error()
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now()
			updated = true
		}
	}

	if updated {
		err = WriteFile(tasks, filename)
		if err != nil {
			return false, err.Error()
		}
	} else {
		return false, "Task with that ID not found"
	}

	return true, ""
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	descWidth := 45
	statusWidth := 12

	fmt.Printf(" %-4s | %-*s | %-*s | %-19s | %-19s\n",
		"ID", descWidth, "Description", statusWidth, "Status", "Updated at", "Created at")

	fmt.Println(strings.Repeat("-", 4+3+descWidth+3+statusWidth+3+19+3+19))

	for _, task := range tasks {
		desc := task.Description
		if len(desc) > descWidth {
			desc = desc[:descWidth-3] + "..."
		}

		fmt.Printf(" %-4d | %-*s | %-*s | %s | %s\n",
			task.ID,
			descWidth, desc,
			statusWidth, task.Status,
			task.UpdatedAt.Format("2006-01-02 15:04:05"),
			task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
