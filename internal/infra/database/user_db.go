package database

import (
	"database/sql"
	"github.com/MatThHeuss/si_2020_2_api/internal/dto"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"log"
)

type User struct {
	DB *sql.DB
}

func NewUserDb(db *sql.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {

	stmt, err := u.DB.Prepare("insert into users (id, name, profile_image_url, email, password) values (?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error in prepare statement: %s", err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.ProfileImageURL, user.Email, user.Password)

	if err != nil {
		log.Printf("Error im exec query: %s", err)
		return err
	}

	log.Println("Success in insertion")
	return nil
}

func (u *User) FindByEmail(email string) (*dto.CreateUserOutput, error) {
	stmt, err := u.DB.Prepare("select id, name, email, profile_image_url, created_at from users where email = ?")
	if err != nil {
		log.Printf("Error in prepare statement: %s", err)
		return nil, err
	}
	defer stmt.Close()

	var user dto.CreateUserOutput
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.ProfileImageURL, &user.CreatedAt)
	if err != nil {
		log.Printf("Error in Scan: %s", err)
		return nil, err
	}
	return &user, nil
}
