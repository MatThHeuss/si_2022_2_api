package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/MatThHeuss/si_2020_2_api/internal/dto"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/database"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/gcp"
	"log"
	"net/http"
)

type UserHandler struct {
	UserDb database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{db}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user dto.CreateUserInput
	r.ParseMultipartForm(100)
	mForm := r.MultipartForm

	for k, _ := range mForm.File {
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Println("inovke FormFile error:", err)
			return
		}

		err = gcp.UploadFile(fileHeader.Filename, file)
		if err != nil {
			log.Println(err)
		}

		photoUrl := fmt.Sprintf("https://storage.googleapis.com/si_images_unb/%s", fileHeader.Filename)
		user.ProfileImageURL = photoUrl
		file.Close()
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password, user.ProfileImageURL)

	fmt.Println(u)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDb.Create(u)
	if err != nil {
		log.Printf("Error inserting user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}
