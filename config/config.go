package config

const BotToken string = "5610958233:AAF7iMAJBZdrEJdJsQF9GJ_Oxcm7S_TiNio"

const YandexTocken string = "25b77834-9b16-43b1-8b8e-28c2742f4819"
const YandexUrl string = "https://api.weather.yandex.ru/v2/forecast" //?lat=45.043317&lon=41.969110"
const YandexKey string = "X-Yandex-API-Key"

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
