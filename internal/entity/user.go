package entity

import (
	"log"

	"github.com/MatThHeuss/si_2020_2_api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           entity.ID `json:"id"`
	Name         string    `json:"name"`
	ProfileImage string    `json:"profile_image"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
}

func NewUser(name, email, password string, profile_image string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("aqui entao")
		return nil, err
	}

	user := &User{
		ID:           entity.NewID(),
		Name:         name,
		ProfileImage: profile_image,
		Email:        email,
		Password:     string(hash),
	}

	// err = user.Validate()

	// if err != nil {
	// 	return nil, err
	// }
	return user, nil
}
