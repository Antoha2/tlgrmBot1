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

	tx := r.rep.Table("messagelist").Where("chat_id = ?", ChatId).Last(ms).Scan(ms)

	if tx.RowsAffected == 0 {
		log.Println("записей нет")
	}
	return ms, nil
}

func (r *repositoryImplDB) AddUser(user *RepositoryUserlist) error {

	result := r.rep.Table("userlist").Create(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка добавления пользователя")
	}
	log.Println("добавлен пользователь с id - ", user.UserId)
	return nil
	// log.Println(user)

}

func (r *repositoryImplDB) UserVerification(user *RepositoryUserlist) bool {

	var count int64
	r.rep.Table("userlist").Where("user_id = ?", user.UserId).Find(&user).Count(&count)
	if count == 0 {
		log.Printf("пользователь с Id %d не найден\n", user.UserId)
		return false
	}
	log.Printf("пользователь с Id %d найден\n", user.UserId)
	return true
}
