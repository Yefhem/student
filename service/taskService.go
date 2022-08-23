package service

import (
	"strconv"

	"github.com/Yefhem/student-syllabus/model"
	"github.com/Yefhem/student-syllabus/repository"
)

type TaskService interface {
	CreateTask(task model.TaskDTO) error
	UpdateTask(taskDTO model.TaskDTO, taskID string) error
	GetTask(taskID string) (model.Task, error)
	DeleteTask(taskID string) error
	GetAllTasks() ([]model.Task, error)

	StatusService(status string, taskID string) error
	TotalTask() (int64, error)
	LastTask() (model.Task, error)
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

// --------------------> Methods...
// ---------->
func (t *taskService) CreateTask(taskDTO model.TaskDTO) error {

	statem := model.State{
		Canceled:  false,
		Finished:  false,
		Continues: true,
	}

	task := model.Task{
		Title:       taskDTO.Title,
		Description: taskDTO.Description,
		Date:        taskDTO.Date,
		Rate:        taskDTO.Rate,
		State:       statem,
	}

	if err := t.taskRepo.Create(task); err != nil {
		return err
	}

	return nil
}

func (t *taskService) UpdateTask(taskDTO model.TaskDTO, taskID string) error {

	currentTask, _ := t.GetTask(taskID)

	if currentTask.Title != taskDTO.Title {
		currentTask.Title = taskDTO.Title
	}
	if currentTask.Description != taskDTO.Description {
		currentTask.Description = taskDTO.Description
	}
	if currentTask.Date != taskDTO.Date {
		currentTask.Date = taskDTO.Date
	}
	if currentTask.Rate != taskDTO.Rate {
		currentTask.Rate = taskDTO.Rate
	}

	if err := t.taskRepo.Update(currentTask); err != nil {
		return err
	}

	return nil
}

// ---------->
func (t *taskService) GetTask(taskID string) (model.Task, error) {

	convID, _ := strconv.ParseUint(taskID, 10, 64)

	task, err := t.taskRepo.Get(convID)
	if err != nil {
		return task, err
	}
	return task, nil
}

// ---------->
func (t *taskService) DeleteTask(taskID string) error {

	convID, _ := strconv.ParseUint(taskID, 10, 64)

	if err := t.taskRepo.Delete(convID); err != nil {
		return err
	}

	return nil
}

// ---------->
func (t *taskService) GetAllTasks() ([]model.Task, error) {
	tasks, err := t.taskRepo.GetAll()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (t *taskService) StatusService(status string, taskID string) error {

	var state = model.State{}

	if status == "completed" {
		state.Finished = true
	} else {
		state.Canceled = true
	}

	currentTask, _ := t.GetTask(taskID)

	currentTask.State = state

	if err := t.taskRepo.Update(currentTask); err != nil {
		return err
	}

	return nil
}

func (t *taskService) TotalTask() (int64, error) {
	count, err := t.taskRepo.TotalTasks()
	if err != nil {
		return count, err
	}

	return count, nil
}

func (t *taskService) LastTask() (model.Task, error) {

	lastTask, err := t.taskRepo.LastTask()
	if err != nil {
		return lastTask, err
	}

	return lastTask, nil
}
