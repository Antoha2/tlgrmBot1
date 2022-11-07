package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (r *repositoryImplDB) AddMessage(ms *RepositoryMessage) error {

	query := "INSERT INTO messagelist (id_chat, user_name, text, response) VALUES ($1, $2, $3, $4) RETURNING id"
	result := r.rep.Table("messagelist").Raw(query, ms.Chat.ChatId, ms.UserName, ms.Text, ms.Response).Scan(&ms.MessageId)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка сознания задачи")
	}
	log.Println("создана запись - ", ms)

	return nil
}
func (r *repositoryImplDB) RepeatMessage(ChatId int64) (*RepositoryMessage, error) {
	result := new(RepositoryMessage)
	query := "SELECT * FROM messagelist WHERE id_chat = $1 ORDER BY id DESC LIMIT 1"
	tx := r.rep.Table("messagelist").Raw(query, ChatId).Scan(result)
	if tx.RowsAffected == 0 {
		log.Println("записей нет")
	}
	return result, nil
}
