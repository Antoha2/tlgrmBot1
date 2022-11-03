package meteo

type GetWinder interface {
	GetWind(apiUrl, apiKey, apiTocken string) (string, error)
}
