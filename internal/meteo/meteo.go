package meteo

type GetWinder interface {
	GetWind(request *Querry) (string, error)
}

type Querry struct {
	Lat      string
	Lon      string
	CityName string
}
