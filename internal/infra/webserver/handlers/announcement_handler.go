package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/database"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/gcp"
	pkg "github.com/MatThHeuss/si_2020_2_api/pkg/entity"
	"log"
	"net/http"
	"strings"
)

type AnnouncementHandler struct {
	AnnouncementDb      database.AnnouncementInterface
	AnnouncementImageDb database.AnnouncementImagesInterface
}

func NewAnnouncementHandler(announcementDb database.AnnouncementInterface, announcementImageDb database.AnnouncementImagesInterface) *AnnouncementHandler {
	return &AnnouncementHandler{announcementDb, announcementImageDb}
}

func (h *AnnouncementHandler) CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
	var photoUrl string
	r.ParseMultipartForm(1000)
	mForm := r.MultipartForm

	name := mForm.Value["name"]
	description := mForm.Value["description"]
	address := mForm.Value["address"]
	postalCode := mForm.Value["postal_code"]
	userId, err := pkg.ParseID(mForm.Value["user_id"][0])
	if err != nil {
		log.Printf("Error parsing id: %s", err)
	}

	announcement, err := entity.NewAnnouncement(name[0], description[0], address[0], postalCode[0], userId)
	if err != nil {
		log.Printf("Error creating new announcement entity %s", err)
	}

	err = h.AnnouncementDb.Create(announcement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error inserting in announcement table: %s", err)
		return
	}

	for k, _ := range mForm.File {
		file, fileHeader, err := r.FormFile(k)
		fileName := strings.ReplaceAll(fileHeader.Filename, " ", "_")
		if err != nil {
			fmt.Println("inovke FormFile error:", err)
			return
		}
		defer file.Close()
		err = gcp.UploadFile(fileName, file)
		if err != nil {
			log.Println(err)
		}

		photoUrl = fmt.Sprintf("https://storage.googleapis.com/si_images_unb/%s", fileName)
		announcementImage, err := entity.NewAnnouncementImage(announcement.ID.String(), photoUrl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("error creating new announcement image entity: %s", err)
			return
		}
		err = h.AnnouncementImageDb.Create(announcementImage)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("error inserting in announcement image table: %s", err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AnnouncementHandler) GetAllAnnouncements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	announcements, err := h.AnnouncementDb.GetAllAnnouncements()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error loading all announcements: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(announcements)
}
