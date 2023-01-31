package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"github.com/MatThHeuss/si_2020_2_api/internal/errors"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/database"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/gcp"
	pkg "github.com/MatThHeuss/si_2020_2_api/pkg/entity"
	"github.com/go-chi/chi"
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
	category := mForm.Value["category"]
	description := mForm.Value["description"]
	address := mForm.Value["address"]
	postalCode := mForm.Value["postal_code"]
	userId, err := pkg.ParseID(mForm.Value["user_id"][0])
	if err != nil {
		log.Printf("Error parsing id: %s", err)
	}

	announcement, err := entity.NewAnnouncement(name[0], category[0], description[0], address[0], postalCode[0], userId)
	if err != nil {
		log.Printf("Error creating new announcement entity %s", err)
	}

	err = h.AnnouncementDb.Create(announcement)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error inserting in announcement table: %s", err)
		return
	}

	for k, _ := range mForm.File {
		file, fileHeader, err := r.FormFile(k)
		fileName := strings.ReplaceAll(fileHeader.Filename, " ", "_")
		if err != nil {
			log.Printf("inovke FormFile error: %s", err)
			return
		}

		err = gcp.UploadFile(fileName, file)
		if err != nil {
			log.Printf("Error uploading file: %s", err)
		}

		photoUrl = fmt.Sprintf("https://storage.googleapis.com/si_images_unb/%s", fileName)
		announcementImage, err := entity.NewAnnouncementImage(announcement.ID.String(), photoUrl)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("error creating new announcement image entity: %s", err)
			return
		}
		err = h.AnnouncementImageDb.Create(announcementImage)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("error inserting in announcement image table: %s", err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AnnouncementHandler) GetAllAnnouncements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	announcements, err := h.AnnouncementDb.GetAllAnnouncements()
	if err != nil {
		err := errors.Errors{
			Message:    "No one announcement found",
			StatusCode: http.StatusNotFound,
		}
		w.WriteHeader(http.StatusNotFound)
		log.Printf("error loading all announcements: %s", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(announcements)
}

func (h *AnnouncementHandler) GetAnnouncementById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	announcements, err := h.AnnouncementDb.GetAnnouncementById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := &errors.Errors{
			Message:    "announcement not found",
			StatusCode: http.StatusNotFound,
		}
		json.NewEncoder(w).Encode(err)
		log.Printf("error loading  announcement: %s", err)
		return
	}

	if announcements == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("announcement not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(announcements)
}

func (h *AnnouncementHandler) GetAnnouncementByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	category := chi.URLParam(r, "category")
	announcements, err := h.AnnouncementDb.GetAnnouncementByCategory(category)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := &errors.Errors{
			Message:    "announcement not found",
			StatusCode: http.StatusNotFound,
		}
		json.NewEncoder(w).Encode(err)
		log.Printf("error loading  announcement: %s", err)
		return
	}

	if announcements == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("announcement not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(announcements)
}
