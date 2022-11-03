package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (r *repositoryImplDB) AddMessage(ms *RepositoryMessage) error {

	query := "INSERT INTO messagelist (id_chat, user_name, text) VALUES ($1, $2, $3) RETURNING id"
	result := r.rep.Table("messagelist").Raw(query, ms.Chat.ChatId, ms.UserName, ms.Text).Scan(&ms.MessageId)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка сознания задачи")
	}
	log.Println("создана запись - ", ms)

	return nil
}
