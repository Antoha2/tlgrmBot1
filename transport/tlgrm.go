package transport

const botToken string = "5610958233:AAF7iMAJBZdrEJdJsQF9GJ_Oxcm7S_TiNio"

const GismeteoToken string = "56b30cb255.3443075"
const GismeteoApi string = "https://api.gismeteo.net/v2/search/cities/?query=москв"

const YandexTocken string = "25b77834-9b16-43b1-8b8e-28c2742f4819"

//const botApi string = "https://api.telegram.org/bot"

type Gismeteo struct {
	Wind GisWind `json:"wind"`
	//Message  Message `json:"message"`
}

type GisWind struct {
	Direction GisWindDirection `json:"direction"`
	Speed     GisWindSpeed     `json:"speed"`
}

type GisWindDirection struct {
	Scale_8 int `json:"scale_8"`
	Degree  int `json:"degree"`
}

type GisWindSpeed struct {
	M_s float32 `json:"m_s"`
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

/*
type Gismeteo struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
} */
