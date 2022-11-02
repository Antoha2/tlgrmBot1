package config

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	Sslmode  string
}

func GetConfig() *Config {

	return &Config{
		DB: DBConfig{
			Dbname:   "bot",
			User:     "user",
			Password: "user",
			Host:     "postgres",
			Port:     5432,
			Sslmode:  "",
		},
	}

}

/* DB: DBConfig{
	Dbname:   "root",
	User:     "root",
	Password: "root",
	Host:     "postgres",
	Port:     5432,
	Sslmode:  "",
}, */
