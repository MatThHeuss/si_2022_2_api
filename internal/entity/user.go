package entity

import (
	"github.com/MatThHeuss/si_2020_2_api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID              entity.ID `json:"id"`
	Name            string    `json:"name"`
	ProfileImageURL string    `json:"profile_image_url"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewUser(name, email, password string, profileImageURL string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:              entity.NewID(),
		Name:            name,
		ProfileImageURL: profileImageURL,
		Email:           email,
		Password:        string(hash),
		CreatedAt:       time.Now(),
	}

	// err = user.Validate()

	// if err != nil {
	// 	return nil, err
	// }
	return user, nil
}
