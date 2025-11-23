package repository

import (
	"errors"
	"livecode/internal/domain/entity"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	FindByID(int) (*entity.User, error)
	BulkInsert(users []entity.User) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindByID(id int) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("id = ?", id).Preload("Addresses").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *userRepository) BulkInsert(users []entity.User) error {
	return u.db.CreateInBatches(&users, 100).Error
}
