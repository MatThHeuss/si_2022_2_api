package entity

import "github.com/MatThHeuss/si_2020_2_api/pkg/entity"

type Product struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserId      entity.ID `json:"user_id"`
}

func NewProduct(name string, description string, userId entity.ID) (*Product, error) {
	product := &Product{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		UserId:      userId,
	}

	return product, nil
}
