# go-tasks

CLI task manager written in go.

- Project idea from [roadmap.sh](https://roadmap.sh/projects/task-tracker)

### Available commands

```bash
# Add a new task
go-tasks add "task description"

# Update the description of a task
go-tasks update 3 "new description"

# Delete a task
go-tasks delete 5

# Mark a task as in progress
go-tasks mark-in-progress 5

# Mark a task as done
go-tasks mark-done 5

# List tasks
go-tasks list
go-tasks list done
go-tasks list todo
go-tasks list in-progress
```
