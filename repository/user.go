package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	// Query data user by email
	data := r.db.Where("email = ?",
		email).Find(&user)
	// Error Handling
	if data.Error != nil {
		if data.Error == gorm.ErrRecordNotFound {
			return model.User{},
				nil
		}
		return model.User{},
			data.Error
	}

	return user,
		nil // TODO: replace this
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var userTaskCategory []model.UserTaskCategory
	// Joining table
	data := r.db.Table("users").Select("users.id as ID, users.fullname as Fullname, users.email as Email, tasks.title as Task, tasks.deadline as Deadline, tasks.priority as Priority, tasks.status as Status, categories.name as Category").Joins(
		"left join tasks on tasks.user_id = users.id").Joins(
		"left join categories on categories.id = tasks.category_id").Scan(&userTaskCategory)
	if data.Error != nil {
		return []model.UserTaskCategory{},
			data.Error
	}

	return userTaskCategory, nil // TODO: replace this
}
