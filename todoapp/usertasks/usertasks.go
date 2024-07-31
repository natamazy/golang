package usertasks

import (
	"errors"
	"fmt"

	"todo-app.com/task"
)

type UserTasks struct {
	Username              string      `json:"username"`
	Tasks                 []task.Task `json:"tasks"`
	TotalTasksCount       int         `json:"total_tasks_count"`
	CompletedTasksCount   int         `json:"completed_tasks_count"`
	NotFinishedTasksCount int         `json:"not_finished_tasks_count"`
}

func New(username string) (*UserTasks, error) {
	if username == "" {
		return nil, errors.New("username can't be empty")
	}

	return &UserTasks{
		Username:              username,
		Tasks:                 make([]task.Task, 0),
		TotalTasksCount:       0,
		CompletedTasksCount:   0,
		NotFinishedTasksCount: 0,
	}, nil
}

func (userTasks *UserTasks) Add(task task.Task) error {
	if task.TaskName == "" || task.TaskDescription == "" {
		return errors.New("task name or description can't be empty")
	}

	userTasks.Tasks = append(userTasks.Tasks, task)
	return nil
}

func (userTasks *UserTasks) Remove(indexOfTask int) error {
	for index := range userTasks.Tasks {
		if index == indexOfTask {
			userTasks.Tasks = append(userTasks.Tasks[:index], userTasks.Tasks[index+1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}

func (userTasks *UserTasks) ListTasks() error {
	if userTasks.TotalTasksCount == 0 {
		return errors.New("there is no tasks")
	}

	for index, value := range userTasks.Tasks {
		fmt.Printf("%d - Name: %s, Description: %s\n", index+1, value.TaskName, value.TaskDescription)
	}
	return nil
}
