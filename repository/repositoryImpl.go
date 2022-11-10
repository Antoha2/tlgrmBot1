package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

//добавить сообщение в БД
func (r *repositoryImplDB) AddMessage(ms *RepositoryMessagelist) error {

	result := r.rep.Table("messagelist").Create(ms).Scan(&ms)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка сознания задачи")
	}
	log.Println("создана запись - ", ms.MessageId)
	return nil
}

//повтор крайнего запроса
func (r *repositoryImplDB) RepeatMessage(ChatId int64) (*RepositoryMessagelist, error) {
	ms := &RepositoryMessagelist{}

	//query := "SELECT * FROM messagelist WHERE chat_id = $1 ORDER BY id DESC LIMIT 1"
	//tx := r.rep.Table("messagelist").Raw(query, ChatId).Scan(result)

	err := r.rep.Table("messagelist").Where("chat_id = ?", ChatId).Last(ms).Scan(ms).Error
	if err != nil {
		return nil, err
	}

	if ms.ChatId == 0 {
		log.Println("записей нет")
	}
	return ms, nil
}

//добавить в БД нового пользователя
func (r *repositoryImplDB) AddUser(user *RepositoryUserlist) error {

	result := r.rep.Table("userlist").Create(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка добавления пользователя")
	}
	log.Println("добавлен пользователь с id - ", user.UserId)
	return nil
}

//проверка пользователя
func (r *repositoryImplDB) UserVerification(user *RepositoryUserlist) bool {

	var count int64
	r.rep.Table("userlist").Where("user_id = ?", user.UserId).Find(&user).Count(&count)
	if count == 0 {
		log.Printf("пользователь с Id %d не найден\n", user.UserId)
		return false
	}
	log.Printf("1 пользователь с Id %d найден\n", user.UserId)
	return true
}

//получить историю пользователя
func (r *repositoryImplDB) GetHistory(userId int) ([]*RepositoryMessagelist, error) {

	sliceMsg := make([]*RepositoryMessagelist, 0)
	err := r.rep.Table("messagelist").Where("user_id = ?", userId).Find(&sliceMsg).Scan(&sliceMsg).Error
	if err != nil {
		log.Println("GetHistory(rep) сообщения не найдены - ", err)
		return nil, err
	}
	return sliceMsg, nil
}
