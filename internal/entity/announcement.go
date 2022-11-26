package entity

import "github.com/MatThHeuss/si_2020_2_api/pkg/entity"

type Announcement struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	PostalCode  string    `json:"postal_code"`
	UserId      string    `json:"user_id"`
}

func NewAnnouncement(name string, description string, address string, postalCode string, userId entity.ID) (*Announcement, error) {

	announcement := &Announcement{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		Address:     address,
		PostalCode:  postalCode,
		UserId:      userId.String(),
	}

	return announcement, nil
}
