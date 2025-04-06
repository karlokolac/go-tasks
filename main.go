package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const filename = "~/Documents/tasks.json"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No commands were given. Use 'task-cli help' to list all commands")
	}

	command := os.Args[1]
	switch command {
	case "add":
		if CheckArgsLength(os.Args, 3) {
			ok, str := AddTask(os.Args[2], filename)
			if !ok {
				log.Fatal(str)
			}
			log.Println("Task added successfully")
		} else {
			log.Fatal("Wrong number of arguments for add command.")
		}

	case "update":
		if CheckArgsLength(os.Args, 4) {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}

			ok, str := UpdateTask(id, os.Args[3], filename)
			if !ok {
				log.Fatal(str)
			}
			log.Println("Task updated successfully")
		} else {
			log.Fatal("Wrong number of arguments for update command")
		}

	case "delete":
		if CheckArgsLength(os.Args, 3) {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}

			ok, str := DeleteTask(id, filename)
			if !ok {
				log.Fatal(str)
			}
			log.Println("Task deleted successfully")
		} else {
			log.Fatal("Wrong number of arguments for delete command")
		}

	case "mark-in-progress":
		if CheckArgsLength(os.Args, 3) {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}

			ok, str := MarkTaskInProgress(id, filename)
			if !ok {
				log.Fatal(str)
			}
			log.Println("Task updated successfully")
		} else {
			log.Fatal("Wrong number of arguments for mark-in-progress command")
		}

	case "mark-done":
		if CheckArgsLength(os.Args, 3) {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}

			ok, str := MarkTaskDone(id, filename)
			if !ok {
				log.Fatal(str)
			}
			log.Println("Task updated successfully")
		} else {
			log.Fatal("Wrong number of arguments for mark-done command")
		}

	case "list": // [done | todo | in-progress]
		if CheckArgsLength(os.Args, 2) {
			tasks, err := ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			ListTasks(tasks)
		} else if CheckArgsLength(os.Args, 3) && os.Args[2] == "in-progress" {
			tasks, err := ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}

			var InProgressTasks []Task
			for _, task := range tasks {
				if task.Status == "in progress" {
					InProgressTasks = append(InProgressTasks, task)
				}
			}
			ListTasks(InProgressTasks)
		} else if CheckArgsLength(os.Args, 3) && os.Args[2] == "done" {
			tasks, err := ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}

			var DoneTasks []Task
			for _, task := range tasks {
				if task.Status == "done" {
					DoneTasks = append(DoneTasks, task)
				}
			}
			ListTasks(DoneTasks)
		} else if CheckArgsLength(os.Args, 3) && os.Args[2] == "todo" {
			tasks, err := ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}

			var TodoTasks []Task
			for _, task := range tasks {
				if task.Status == "todo" {
					TodoTasks = append(TodoTasks, task)
				}
			}
			ListTasks(TodoTasks)
		} else {
			log.Fatal("Wrong number of arguments for list command or incorrect subcommand")
		}

	case "help":
		fmt.Print(`
		List of all commands:

			task-cli add <string> 		- Add a new task to the list
			task-cli update <int> <string> 	- Update the description of a task
			task-cli delete <int> 		- Delete a task

			task-cli mark-in-progress <int> - Mark task status "in progress"
			task-cli mark-done <int> 	- Mark task status "done"

			task-cli list 			- List all tasks
			task-cli list done 		- List all tasks with status "done"
			task-cli list todo 		- List all tasks with status "todo"
			task-cli list in-progress 	- List all tasks with status "in progress"`)
		fmt.Println()
	}
}
