package task

import (
	"errors"
	"time"
)

// ImportTaks(fileName string) error
// ExportTasks() error

type Task struct {
	TaskName        string    `json:"task_name"`
	TaskDescription string    `json:"task_desc"`
	CreatedAt       time.Time `json:"created_at"`
	Deadline        time.Time `json:"deadline"`
	Completed       bool      `json:"completed"`
	CompilationDate time.Time `json:"compilation_date"`
}

func New(taskName string, taskDescription string, deadline string) (*Task, error) {
	if taskName == "" || taskDescription == "" || deadline == "" {
		return nil, errors.New("name or description can't be empty")
	}

	return &Task{
		TaskName:        taskName,
		TaskDescription: taskDescription,
		CreatedAt:       time.Now(),
		Deadline:        time.Now(),
		Completed:       false,
		CompilationDate: time.Time{},
	}, nil
}

func (task *Task) ChangeName(newName string) error {
	if newName == "" {
		return errors.New("name can't be empty")
	}

	task.TaskName = newName
	return nil
}

func (task *Task) ChangeDescription(newDescription string) error {
	if newDescription == "" {
		return errors.New("description can't be empty")
	}

	task.TaskDescription = newDescription
	return nil
}

func (task *Task) MarkAsDone() error {
	if task.Completed {
		return errors.New("it's already complited")
	}

	task.Completed = true
	return nil
}

func (task *Task) MarkAsUndone() error {
	if !task.Completed {
		return errors.New("it's already uncomplited")
	}

	task.Completed = false
	return nil
}
