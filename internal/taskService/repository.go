package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id string, task Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию
func (r *taskRepository) CreateTask(task Task) (Task, error) {

	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}

	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id string, task Task) (Task, error) {
	result := r.db.First(&task, id)
	if result.Error != nil {
		return Task{}, result.Error
	}

	if task.IsDone == false {
		task.IsDone = true
	} else {
		task.IsDone = false
	}

	result = r.db.Save(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}

	return task, nil
}

func (r *taskRepository) DeleteTaskByID(id string) error {
	var task Task
	err := r.db.Delete(&task, id).Error
	return err
}
