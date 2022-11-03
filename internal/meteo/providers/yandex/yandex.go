package yandex

import "github.com/Antoha2/tlgrmBot1/internal/meteo"

const YandexTocken string = "25b77834-9b16-43b1-8b8e-28c2742f4819"
const YandexUrl string = "https://api.weather.yandex.ru/v2/forecast?lat=45.043317&lon=41.969110"
const YandexKey string = "X-Yandex-API-Key"

type yandexImpl struct {
	meteo.GetWinder
}

func NewYandex() *yandexImpl {
	return &yandexImpl{}
}

type Yandex struct {
	Now_dt     string          `json:"now_dt"`
	Info       YandexInfo      `json:"info"`
	Geo_object YandexGeoObject `json:"geo_object"`
	Fact       YandexFact      `json:"fact"`
	//forecasts  YandeFforecasts `json:"forecasts"`
}

type YandexInfo struct {
	Url string `json:"url"`
}
type YandexGeoObject struct {
	Locality YandexLocality `json:"locality"`
	Province YandexProvince `json:"province"`
}

type YandexLocality struct {
	Name string `json:"name"`
}
type YandexProvince struct {
	Name string `json:"name"`
}
type YandexFact struct {
	Temp       int     `json:"temp"`
	Feels_like int     `json:"feels_like"`
	Condition  string  `json:"condition"`
	Wind_speed float32 `json:"wind_speed"`
	Wind_gust  float32 `json:"wind_gust"`
	Wind_dir   string  `json:"wind_dir"`
}
