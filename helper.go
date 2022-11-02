package tlgrmbot1

type Message struct {
	MessageId int    `json:"id" gorm:"primaryKey`
	UserName  string `json:"user_name" gorm:"column:user_name`
	ChatId    int64  `json:"chat" gorm:"column:id_chat`
	Text      string `json:"text" gorm:"column:text`
}

/* type chat struct {
	ChatId int64 `json:"id"`
}
*/
