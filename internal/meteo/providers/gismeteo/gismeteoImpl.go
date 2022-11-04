package gismeteo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Antoha2/tlgrmBot1/config"
	"github.com/Antoha2/tlgrmBot1/internal/meteo"
)

func (s *gismeteoImpl) GetWind(request *meteo.Querry) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", config.GismeteoUrl, nil)
	if err != nil {
		log.Println("http.NewRequest() - ", err)
		return "", err
	}
	req.Header.Set("apiKey", config.GismeteoToken) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! apiKey
	log.Println("req - ", req)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do() - ", err)
		return "", err
	}

	defer resp.Body.Close()

	log.Println("resp - ", resp)
	//resp.Header.Set()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err)
		return "", err
	}

	restResponse := new(Gismeteo)
	err = json.Unmarshal(body, restResponse)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return "", err
	}
	log.Println(restResponse)
	//msg:=string(restResponse)
	return "", nil

}
