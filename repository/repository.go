package repository

import "gorm.io/gorm"

type Repository interface {
	AddMessage(*RepositoryMessage) error
	RepeatMessage(ChatId int64) (*RepositoryMessage, error)
}

type repositoryImplDB struct {
	rep *gorm.DB
	Repository
}

func NewRepository(dbx *gorm.DB) *repositoryImplDB {
	return &repositoryImplDB{
		rep: dbx,
	}
}

type RepositoryMessage struct {
	MessageId int    `json:"id"`
	UserName  string `json:"user_name"`
	Chat      chat   `json:"chat"`
	Text      string `json:"text"`
	Response  string `json:"response"`
}

type chat struct {
	ChatId int64 `json:"chat_id"`
}
