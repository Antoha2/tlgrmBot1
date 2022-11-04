package yandex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Antoha2/tlgrmBot1/config"
	"github.com/Antoha2/tlgrmBot1/internal/meteo"
)

func (s *yandexImpl) GetWind(request *meteo.Querry) (string, error) {

	client := &http.Client{}
	//log.Println(geokoder.Coordinates)
	req, err := http.NewRequest("GET", config.YandexUrl, nil)
	if err != nil {
		log.Println("http.NewRequest() - ", err)
		return "", err
	}
	req.Header.Set(config.YandexKey, config.YandexTocken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do() - ", err)
		return "", err
	}
	defer resp.Body.Close()

	// log.Println("resp - ", resp)
	//resp.Header.Set()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err)
		return "", err
	}

	r := new(Yandex)
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return "", err
	}

	var wind_dir string
	switch r.Fact.Wind_dir {
	case "nw":
		wind_dir = "северо-западное"
	case "n":
		wind_dir = "северное"
	case "ne":
		wind_dir = "северо-восточное"
	case "e":
		wind_dir = "восточное"
	case "se":
		wind_dir = "юго-восточное"
	case "s":
		wind_dir = "южное"
	case "sw":
		wind_dir = "юго-западное"
	case "w":
		wind_dir = "западное"
	case "c":
		wind_dir = "штиль"

	}

	msg := fmt.Sprintf("Яндекс \n %s %s \n Скорость ветра (в м/с) - %.1f\n Скорость порывов ветра (в м/с) - %.1f\n Направление ветра - %s",
		r.Geo_object.Locality.Name, r.Geo_object.Province.Name, r.Fact.Wind_speed, r.Fact.Wind_gust, wind_dir)
	return msg, nil

}

/*

)

/* func getWind(apiUrl, apiKey, apiTocken string) error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Println("http.NewRequest() - ", err)
		return err
	}
	req.Header.Set("X-Yandex-API-Key", YandexTocken)
	log.Println("req - ", req)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do() - ", err)
		return err
	}

	defer resp.Body.Close()

	log.Println("resp - ", resp)
	//resp.Header.Set()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll() -", err)
		return err
	}

	restResponse := new(Yandex)
	err = json.Unmarshal(body, restResponse)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return err
	}
	log.Println(restResponse)
	return nil
}
*/
