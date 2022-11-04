package config

const YandexTocken string = "25b77834-9b16-43b1-8b8e-28c2742f4819"
const YandexUrl string = "https://api.weather.yandex.ru/v2/forecast" //?lat=45.043317&lon=41.969110"
const YandexKey string = "X-Yandex-API-Key"

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
