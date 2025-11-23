package usecase

import (
	"errors"
	"fmt"
	"livecode/internal/adapter/repository"
	"livecode/internal/domain/entity"

	"gorm.io/gorm"
)

type UserService interface {
	GetByID(userid int) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) GetByID(userid int) (*entity.User, error) {
	user, err := u.userRepo.FindByID(userid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	return user, nil
}
