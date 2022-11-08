package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

func (r *repositoryImplDB) AddMessage(ms *RepositoryMessagelist) error {

	result := r.rep.Table("messagelist").Create(ms).Scan(&ms)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка сознания задачи")
	}
	log.Println("создана запись - ", ms.MessageId)

	return nil
}
func (r *repositoryImplDB) RepeatMessage(ChatId int64) (*RepositoryMessagelist, error) {
	ms := &RepositoryMessagelist{}

	//query := "SELECT * FROM messagelist WHERE chat_id = $1 ORDER BY id DESC LIMIT 1"
	//tx := r.rep.Table("messagelist").Raw(query, ChatId).Scan(result)
	//result := map[string]interface{}{}
	//tx := r.rep.Model(&RepositoryMessage{}).Last(&result).Scan(ms)

	tx := r.rep.Table("messagelist").Where("chat_id = ?", ChatId).Last(ms).Scan(ms)

	if tx.RowsAffected == 0 {
		log.Println("записей нет")
	}
	return ms, nil
}
