CREATE TABLE messageList (
	id SERIAL PRIMARY KEY,  
	user_id INT NOT NULL,
	chat_id INT NOT NULL,
	text VARCHAR(255)  NOT NULL,
  	response VARCHAR(255)  NOT NULL,
  	FOREIGN KEY (user_id) REFERENCES userlist(user_id)
)
