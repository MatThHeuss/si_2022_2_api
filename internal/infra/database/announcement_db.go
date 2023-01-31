package database

import (
	"database/sql"
	"github.com/MatThHeuss/si_2020_2_api/internal/dto"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"log"
	"strings"
)

type Announcement struct {
	DB *sql.DB
}

func NewAnnouncementDb(db *sql.DB) *Announcement {
	return &Announcement{DB: db}
}

func (a *Announcement) Create(announcement *entity.Announcement) error {

	stmt, err := a.DB.Prepare("insert into announcement (id, name, category, description, address, postal_code, user_id) VALUES (?,?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error in prepare statement: %s", err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(announcement.ID, announcement.Name, announcement.Category, announcement.Description, announcement.Address, announcement.PostalCode, announcement.UserId)

	if err != nil {
		log.Printf("Error im exec query: %s", err)
		return err
	}

	log.Println("Success in insertion")
	return nil

}

func (a *Announcement) GetAllAnnouncements() (*[]dto.GetAllAnnouncementsOutputToUser, error) {
	//var announcementsOutputs []dto.GetAllAnnouncementsOutput
	var announcementOutputsUsers []dto.GetAllAnnouncementsOutputToUser
	rows, err := a.DB.Query(" select a.id,  a.name, a.category, a.description,  a.address,  a.postal_code,  u.name,  group_concat(image_url) as \"images\" from  users u, announcement a,  announcement_images WHERE  a.user_id = u.id  AND a.id = announcement_images.announcement_id  GROUP BY a.id, a.name, a.category, a.description,a.address,a.postal_code,  u.name;")
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var announcementsOutput dto.GetAllAnnouncementsOutput
		var announcementsOutputUser dto.GetAllAnnouncementsOutputToUser
		err = rows.Scan(&announcementsOutput.ID, &announcementsOutput.Name, &announcementsOutput.Category, &announcementsOutput.Description, &announcementsOutput.Address, &announcementsOutput.PostalCode, &announcementsOutput.User, &announcementsOutput.Images)

		//announcementsOutputs = append(announcementsOutputs, announcementsOutput)
		announcementsOutputUser.ID = announcementsOutput.ID
		announcementsOutputUser.User = announcementsOutput.User
		announcementsOutputUser.Category = announcementsOutput.Category
		announcementsOutputUser.Description = announcementsOutput.Description
		announcementsOutputUser.PostalCode = announcementsOutput.PostalCode
		announcementsOutputUser.Name = announcementsOutput.Name
		announcementsOutputUser.Address = announcementsOutput.Address
		announcementsOutputUser.Images = append(announcementsOutputUser.Images, strings.Split(announcementsOutput.Images, ",")...)

		announcementOutputsUsers = append(announcementOutputsUsers, announcementsOutputUser)

	}

	return &announcementOutputsUsers, nil
}

func (a *Announcement) GetAnnouncementById(id string) (*dto.GetAllAnnouncementsOutputToUser, error) {

	query := " select a.id, a.name,  a.category,  a.description,  a.address,  a.postal_code,  u.name,  group_concat(image_url) as \"images\"\nfrom\n  users u,\n  announcement a,\n  announcement_images\nWHERE\n  a.user_id = u.id\n  AND a.id = announcement_images.announcement_id\n  AND a.id = ?  \n  GROUP BY\n  a.id,\n  a.name,  \n  a.description,\n  a.address,\n  a.postal_code,\n  u.name;"

	var announcementsOutput dto.GetAllAnnouncementsOutput
	if err := a.DB.QueryRow(query, id).Scan(&announcementsOutput.ID, &announcementsOutput.Name, &announcementsOutput.Category, &announcementsOutput.Description, &announcementsOutput.Address, &announcementsOutput.PostalCode, &announcementsOutput.User, &announcementsOutput.Images); err != nil {
		log.Printf("Error executing query: %s", err)
		return nil, err

	}

	var announcementsOutputUser dto.GetAllAnnouncementsOutputToUser

	announcementsOutputUser.ID = announcementsOutput.ID
	announcementsOutputUser.User = announcementsOutput.User
	announcementsOutputUser.Category = announcementsOutput.Category
	announcementsOutputUser.Description = announcementsOutput.Description
	announcementsOutputUser.PostalCode = announcementsOutput.PostalCode
	announcementsOutputUser.Name = announcementsOutput.Name
	announcementsOutputUser.Address = announcementsOutput.Address
	announcementsOutputUser.Images = append(announcementsOutputUser.Images, strings.Split(announcementsOutput.Images, ",")...)
	return &announcementsOutputUser, nil
}

func (a *Announcement) GetAnnouncementByCategory(category string) (*[]dto.GetAllAnnouncementsOutputToUser, error) {
	//var announcementsOutputs []dto.GetAllAnnouncementsOutput
	var announcementOutputsUsers []dto.GetAllAnnouncementsOutputToUser
	rows, err := a.DB.Query(" select a.id,  a.name, a.category, a.description,  a.address,  a.postal_code,  u.name,  group_concat(image_url) as \"images\" from  users u, announcement a,  announcement_images WHERE  a.user_id = u.id  AND a.id = announcement_images.announcement_id AND a.category =?  GROUP BY a.id, a.name, a.category, a.description,a.address,a.postal_code,  u.name;", category)
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var announcementsOutput dto.GetAllAnnouncementsOutput
		var announcementsOutputUser dto.GetAllAnnouncementsOutputToUser
		err = rows.Scan(&announcementsOutput.ID, &announcementsOutput.Name, &announcementsOutput.Category, &announcementsOutput.Description, &announcementsOutput.Address, &announcementsOutput.PostalCode, &announcementsOutput.User, &announcementsOutput.Images)

		//announcementsOutputs = append(announcementsOutputs, announcementsOutput)
		announcementsOutputUser.ID = announcementsOutput.ID
		announcementsOutputUser.User = announcementsOutput.User
		announcementsOutputUser.Category = announcementsOutput.Category
		announcementsOutputUser.Description = announcementsOutput.Description
		announcementsOutputUser.PostalCode = announcementsOutput.PostalCode
		announcementsOutputUser.Name = announcementsOutput.Name
		announcementsOutputUser.Address = announcementsOutput.Address
		announcementsOutputUser.Images = append(announcementsOutputUser.Images, strings.Split(announcementsOutput.Images, ",")...)

		announcementOutputsUsers = append(announcementOutputsUsers, announcementsOutputUser)

	}

	return &announcementOutputsUsers, nil
}
