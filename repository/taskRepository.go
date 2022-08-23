package repository

import (
	"github.com/Yefhem/student-syllabus/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task model.Task) error
	Get(taskID uint64) (model.Task, error)
	GetAll() ([]model.Task, error)
	Update(task model.Task) error
	Delete(taskID uint64) error

	TotalTasks() (int64, error)
	LastTask() (model.Task, error)
}

type taskConnection struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskConnection{
		db: db,
	}
}

// --------------------> Methods...

// ----------> Creates a Task and returns it if there is an error...
func (c *taskConnection) Create(task model.Task) error {

	if err := c.db.Preload("Date").Preload("State").Create(&task).Error; err != nil {
		return err
	}

	return nil
}

// ----------> Find Task By ID
func (c *taskConnection) Get(taskID uint64) (model.Task, error) {

	var task model.Task

	if err := c.db.Preload("Date").Preload("State").First(&task, taskID).Error; err != nil {
		return task, err
	}

	return task, nil
}

// ----------> Get All Tasks
func (c *taskConnection) GetAll() ([]model.Task, error) {
	var tasks []model.Task

	if result := c.db.Preload("Date").Preload("State").Find(&tasks); result.Error != nil {
		return tasks, result.Error
	}

	return tasks, nil
}

// ----------> Update Task By ID
func (c *taskConnection) Update(task model.Task) error {

	if err := c.db.Save(&task).Error; err != nil {
		return err
	}

	return nil
}

// ----------> Delete Task By ID
func (c *taskConnection) Delete(taskID uint64) error {

	task, err := c.Get(taskID)
	if err != nil {
		return err
	}

	if err := c.db.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}

// ----------> Total Task
func (c *taskConnection) TotalTasks() (int64, error) {

	var count int64

	if err := c.db.Table("tasks").Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

// ----------> Last Task
func (c *taskConnection) LastTask() (model.Task, error) {

	var task = model.Task{}

	err := c.db.Last(&task).Error

	// if err := c.db.Last(&task).Error; err != nil {
	// 	return task, err
	// }

	return task, err
}
