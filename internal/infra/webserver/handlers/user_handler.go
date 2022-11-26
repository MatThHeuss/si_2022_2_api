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
	"strings"
)

type UserHandler struct {
	UserDb database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{db}
}

func (h *UserHandler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userdto dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userdto)
	if err != nil {
		log.Printf("Error decoding user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	user, err := h.UserDb.FindByEmail(userdto.Email)
	if err != nil {
		log.Printf("User not found: %s", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1000)
	mForm := r.MultipartForm
	name := mForm.Value["name"]
	email := mForm.Value["email"]
	password := mForm.Value["password"]

	_, err := h.UserDb.FindByEmail(email[0])

	if err == nil {
		log.Println("User already exists")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User already exists")
		return
	}

	file, fileHeader, err := r.FormFile("profile_image")
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

	photoUrl := fmt.Sprintf("https://storage.googleapis.com/si_images_unb/%s", fileName)

	u, err := entity.NewUser(name[0], email[0], password[0], photoUrl)

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
