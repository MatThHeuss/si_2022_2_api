package entity

import "github.com/MatThHeuss/si_2020_2_api/pkg/entity"

type Product struct {
	ID            entity.ID `json:"id"`
	Name          string    `json:"name"`
	ProductImages []string  `json:"product_images"`
	Description   string    `json:"description"`
	UserId        entity.ID `json:"user_id"`
}

func NewProduct(name string, product_images []string, description string, userId entity.ID) (*Product, error) {
	product := &Product{
		Name:          name,
		ProductImages: product_images,
		Description:   description,
		UserId:        userId,
	}

	return product, nil
}
