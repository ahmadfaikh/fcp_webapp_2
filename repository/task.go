package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	err := t.db.Model(model.Task{}).Where("id=?", id).UpdateColumns(map[string]interface{}{
		"title":       task.Title,
		"deadline":    task.Deadline,
		"priority":    task.Priority,
		"category_id": task.CategoryID,
		"status":      task.Status,
	}).Error
	// result := s.db.Save(&model.Student{id, Name: student.Name,})
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t *taskRepository) Delete(id int) error {
	err := t.db.Delete(&model.Task{}, id).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	result := make([]model.Task, 0)
	rows, err := t.db.Table("tasks").Select("*").Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := t.db.ScanRows(rows, &result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil // TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	tasksCategory := make([]model.TaskCategory, 0)
	// rows, err := t.db.Table("tasks").Select("tasks.id as id, tasks.title as title, categories.name as category").Joins("left join categories on tasks.category_id=categories.id").Where("tasks.id = ?", id)
	err := t.db.Model(&tasksCategory).Table("tasks").Select("tasks.id as id, tasks.title as title, categories.name as category").Joins("left join categories on tasks.category_id=categories.id").Where("tasks.id=?", id).Scan(&tasksCategory)
	if err != nil {
		return tasksCategory, err.Error
	}
	return tasksCategory, nil // TODO: replace this
}
