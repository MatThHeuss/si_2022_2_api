package database

import (
	"github.com/MatThHeuss/si_2020_2_api/internal/dto"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*dto.CreateUserOutput, error)
}

type ProductInterface interface {
	Create(user *entity.Product) error
}

type AnnouncementInterface interface {
	Create(announcement *entity.Announcement) error
	GetAllAnnouncements() (*[]dto.GetAllAnnouncementsOutputToUser, error)
	GetAnnouncementById(id string) (*dto.GetAllAnnouncementsOutputToUser, error)
}

type AnnouncementImagesInterface interface {
	Create(user *entity.AnnouncementImage) error
}
