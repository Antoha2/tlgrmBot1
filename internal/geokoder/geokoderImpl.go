package geokoder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *geokoderImpl) GetCoordinates(location string) (*Coordinates, error) {

	locationBody := fmt.Sprintf(`{"query":"г %s"}`, location)
	responseBody := bytes.NewBuffer([]byte(locationBody))
	req, err := http.NewRequest("POST", GeokoderUrl, responseBody)
	if err != nil {
		log.Println("http.NewRequest() - ", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(GeokoderKey, GeokoderTocken)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do() - ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err)
		return nil, err
	}

	r := new(Suggestions)
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return nil, err
	}

	if len(r.Geokoder) == 0 {
		log.Println("координаты не найдены")
		return nil, errors.New("не найдено")
	}

	/* for i, xd := range r.Geokoder {
		log.Println(i, xd)
	} */

	coordinates := new(Coordinates)
	coordinates.Lat = r.Geokoder[0].Data.Geo_lat
	coordinates.Lon = r.Geokoder[0].Data.Geo_lon
	coordinates.CityName = r.Geokoder[0].Unrestricted_value

	/* if coordinates.Lat != "" || coordinates.Lon != "" {
		coordinates.Lat = r.Geokoder[0].Data.Geo_lat
		coordinates.Lon = r.Geokoder[0].Data.Geo_lon
		coordinates.CityName = r.Geokoder[0].Unrestricted_value
	} else {
		coordinates.Lat = r.Geokoder[1].Data.Geo_lat
		coordinates.Lon = r.Geokoder[1].Data.Geo_lon
		coordinates.CityName = r.Geokoder[1].Unrestricted_value
	} */

	return coordinates, nil
}
