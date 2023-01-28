package entity

import (
	"github.com/MatThHeuss/si_2020_2_api/pkg/entity"
	"time"
)

type Chat struct {
	ID         entity.ID `json:"id"`
	SenderID   string    `json:"sender_id"`
	ReceiverID string    `json:"receiver_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"time"`
}

func NewMessage(senderId, receiverId, content string) *Chat {
	chat := &Chat{
		ID:         entity.NewID(),
		SenderID:   senderId,
		ReceiverID: receiverId,
		Content:    content,
		CreatedAt:  time.Now().Local(),
	}

	return chat
}
