package gismeteo

import "github.com/Antoha2/tlgrmBot1/internal/meteo"

const GismeteoToken string = "56b30cb255.3443075"
const GismeteoApi string = "https://api.gismeteo.net/v2/search/cities/?query=москв"

type gismeteoImpl struct {
	meteo.GetWinder
}

func NewGismeteo() *gismeteoImpl {
	return &gismeteoImpl{}
}

type Gismeteo struct {
	Wind GisWind `json:"wind"`
	//Message  Message `json:"message"`GetWinder
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
