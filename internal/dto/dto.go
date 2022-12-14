package dto

import (
	"time"
)

type CreateUserInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ProfileImageURL string `json:"profile_image"`
}

type CreateUserOutput struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	ProfileImageURL string    `json:"profile_image"`
	CreatedAt       time.Time `json:"created_at"`
}

type GetAllAnnouncementsOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	PostalCode  string `json:"postal_code"`
	User        string `json:"user"`
	Images      string `json:"images"`
}

type GetAllAnnouncementsOutputToUser struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Address     string   `json:"address"`
	PostalCode  string   `json:"postal_code"`
	User        string   `json:"user"`
	Images      []string `json:"images"`
}
