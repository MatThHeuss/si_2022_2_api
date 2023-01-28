package handlers

import (
	"encoding/json"
	"github.com/MatThHeuss/si_2020_2_api/internal/dto"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
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
		json.NewEncoder(w).Encode(err)
	}
	c := entity.NewMessage(createChatDto.SenderID, createChatDto.ReceiverID, createChatDto.Content)

	err = h.ChatDb.Create(c)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *ChatHandler) GetAllMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var getAllMessagesDto dto.GetAllMessagesInput
	err := json.NewDecoder(r.Body).Decode(&getAllMessagesDto)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	chats, err := h.ChatDb.GetAllMessages(getAllMessagesDto.SenderID, getAllMessagesDto.ReceiverID)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chats)
}