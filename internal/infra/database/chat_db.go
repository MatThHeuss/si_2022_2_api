package database

import (
	"database/sql"
	"errors"
	"github.com/MatThHeuss/si_2020_2_api/internal/entity"
	"log"
)

type Chat struct {
	DB *sql.DB
}

func NewChatDb(db *sql.DB) *Chat {
	return &Chat{DB: db}
}

func (u *Chat) Create(chat *entity.Chat) error {

	log.Println("Creating chat message")

	stmt, err := u.DB.Prepare("insert into chat (id, sender, receiver, content) values (?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error in prepare statement: %s", err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(chat.ID, chat.SenderID, chat.ReceiverID, chat.Content)

	if err != nil {
		log.Printf("Error im exec query: %s", err)
		return err
	}

	log.Println("Success in insertion")
	return nil
}

func (u *Chat) GetAllMessages(sender, receiver string) (*[]entity.Chat, error) {
	rows, err := u.DB.Query("select * from chat where (sender=? and receiver=?) or (receiver=? and sender=?)", sender, receiver, sender, receiver)
	if err != nil {
		log.Printf("Error im exec query: %s", err)
		return nil, err
	}

	if !rows.Next() {
		log.Printf("Error im exec query: %s", err)
		return nil, errors.New("no message found")
	}

	defer rows.Close()
	var chat entity.Chat
	var chats []entity.Chat
	for rows.Next() {
		err := rows.Scan(&chat.ID, &chat.SenderID, &chat.ReceiverID, &chat.Content, &chat.CreatedAt)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	return &chats, err

}
