package config

import "os"

// const HostAddr = ":8180"
//const ContextKey ContextKey = "History"

const GismeteoToken string = "56b30cb255.3443075"
const GismeteoUrl string = "https://api.gismeteo.net/v2/search/cities/?query=москв"

type Config struct {
	DB   DBConfig
	YA   YAConfig
	TG   TGConfig
	HTTP HTTPConfig
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	Sslmode  string
}

type YAConfig struct {
	YandexTocken string
	YandexUrl    string
	YandexKey    string
}

type TGConfig struct {
	BotToken string
}

type ContextKey string

type HTTPConfig struct {
	HostAddr   string
	ContextKey ContextKey
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
		YA: YAConfig{
			YandexTocken: os.Getenv("YANDEX_TOCKEN"),
			YandexUrl:    os.Getenv("YANDEX_URL"),
			YandexKey:    os.Getenv("YANDEX_KEY"),
		},
		TG: TGConfig{
			BotToken: os.Getenv("BOT_TOCKEN"),
		},
		HTTP: HTTPConfig{
			HostAddr:   os.Getenv("HOST_ADDR"),
			ContextKey: "History",
		},
	}

}
