package config

import "os"

var (
	BotToken     = os.Getenv("BOT_TOCKEN")
	YandexTocken = os.Getenv("YANDEX_TOCKEN")
	YandexUrl    = os.Getenv("YANDEX_URL")
	YandexKey    = os.Getenv("YANDEX_KEY")
)

const GismeteoToken string = "56b30cb255.3443075"
const GismeteoUrl string = "https://api.gismeteo.net/v2/search/cities/?query=москв"

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
