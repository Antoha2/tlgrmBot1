CREATE TABLE userlist (
	id SERIAL,
	user_id INT PRIMARY KEY,
	user_name VARCHAR(255) UNIQUE NOT NULL,
	add_date VARCHAR(255)
)