package handlers

import (
	"encoding/json"
	"github.com/MatThHeuss/si_2020_2_api/internal/dto"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"github.com/MatThHeuss/si_2020_2_api/internal/errors"
	"github.com/MatThHeuss/si_2020_2_api/internal/infra/database"
	"log"
	"net/http"
)

type ChatHandler struct {
	ChatDb database.ChatInterface
}

func NewChatHandler(db database.ChatInterface) *ChatHandler {
	return &ChatHandler{db}
}

func (h *ChatHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createChatDto dto.CreateChatInput
	err := json.NewDecoder(r.Body).Decode(&createChatDto)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	c := entity.NewMessage(createChatDto.SenderID, createChatDto.ReceiverID, createChatDto.Content)

	err = h.ChatDb.Create(c)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *ChatHandler) GetAllMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	senderId := r.URL.Query().Get("sender_id")
	receiverId := r.URL.Query().Get("receiver_id")

	chats, err := h.ChatDb.GetAllMessages(senderId, receiverId)

	if err != nil {
		err := errors.Errors{
			Message:    err.Error(),
			StatusCode: http.StatusNotFound,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chats)
}
