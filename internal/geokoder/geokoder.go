package geokoder

const GeokoderTocken string = "Token 62d34cd3f70c18ff5ead395f3cf9cb5739b7debb"
const GeokoderUrl string = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address"
const GeokoderKey string = "Authorization"

type GeokoderIn interface {
	GetCoordinates(location string) (*Coordinates, error)
}

type geokoderImpl struct {
	GeokoderIn
}

type Coordinates struct {
	Lat      string
	Lon      string
	CityName string
}

func NewGeokoder() *geokoderImpl {
	return &geokoderImpl{}
}

type Suggestions struct {
	Geokoder []Geokoder `json:"suggestions"`
}

type Geokoder struct {
	Value              string `json:"value"`
	Unrestricted_value string `json:"unrestricted_value"`
	Data               Data   `json:"data"`
}

type Data struct {
	Geo_lat string `json:"geo_lat"`
	Geo_lon string `json:"geo_lon"`
}
