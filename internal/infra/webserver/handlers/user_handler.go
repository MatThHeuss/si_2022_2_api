package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MatThHeuss/si_2020_2_api/internal/dto"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password, user.ProfileImage)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(u)

	w.WriteHeader(http.StatusCreated)
}
