package database

import (
	"database/sql"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"log"
)

type AnnouncementImages struct {
	DB *sql.DB
}

func NewAnnouncementImagesDb(db *sql.DB) *AnnouncementImages {
	return &AnnouncementImages{DB: db}
}

func (a *AnnouncementImages) Create(announcementImage *entity.AnnouncementImage) error {

	stmt, err := a.DB.Prepare("insert into announcement_images (announcement_id, image_url) VALUES (?, ?)")

	if err != nil {
		log.Printf("Error in prepare statement: %s", err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(announcementImage.AnnouncementId, announcementImage.ImageURL)

	if err != nil {
		log.Printf("Error im exec query: %s", err)
		return err
	}

	log.Println("Success in insertion")
	return nil

}
