package repository

import "gorm.io/gorm"

type Repository interface {
	AddMessage(*RepositoryMessagelist) error
	RepeatMessage(ChatId int64) (*RepositoryMessagelist, error)
	AddUser(*RepositoryUserlist) error
	UserVerification(*RepositoryUserlist) bool
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

type RepositoryMessagelist struct {
	MessageId int `json:"id" gorm:"column:id"`
	//UserName  string `json:"user_name"`
	UserId   int    `json:"user_id"`
	ChatId   int64  `json:"chat"`
	Text     string `json:"text"`
	Response string `json:"response"`
}

type RepositoryUserlist struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Add_date string `json:"add_date"`
}

/* type RepositoryMessage struct {
	MessageId int    `json:"id"`
	UserName  string `json:"user_name"`
	Chat      chat   `json:"chat"`
	Text      string `json:"text"`
	Response  string `json:"response"`
}

type chat struct {
	ChatId int64 `json:"chat_id"`
} */
