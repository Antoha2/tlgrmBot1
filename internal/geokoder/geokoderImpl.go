package geokoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *geokoderImpl) GetCoordinates(location string) (*Coordinates, error) {

	locationBody := fmt.Sprintf(`{"query":"%s"}`, location) //     `"query/": /"locationBody"`
	//postBody, _ := json.Marshal(`"query/": /"locationBody"`)
	// postBody, _ := json.Marshal(map[string]string{"query": locationBody})
	responseBody := bytes.NewBuffer([]byte(locationBody)) //(postBody)
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

	for i, xd := range r.Geokoder {
		log.Println(i, xd)
	}

	coordinates := new(Coordinates)
	coordinates.Lat = r.Geokoder[0].Data.Geo_lat
	coordinates.Lon = r.Geokoder[0].Data.Geo_lon

	return coordinates, nil
}
