package database

import "github.com/MatThHeuss/si_2020_2_api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
