package yandex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"github.com/Antoha2/tlgrmBot1/config"
	"github.com/Antoha2/tlgrmBot1/internal/meteo"
)

func (s *yandexImpl) GetWind(request *meteo.Querry) (string, error) {

	coordStr := fmt.Sprintf("?lat=%s&lon=%s", request.Lat, request.Lon)

	client := &http.Client{}
	req, err := http.NewRequest("GET", s.config.YA.YandexUrl+coordStr, nil)
	if err != nil {
		log.Println("http.NewRequest() - ", err)
		return "", err
	}
	req.Header.Set(s.config.YA.YandexKey, s.config.YA.YandexTocken)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do() - ", err)
		return "", err
	}
	defer resp.Body.Close()

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

	wind_dir := windDirMap[r.Fact.Wind_dir]

	msg := fmt.Sprintf("Яндекс \n %s %s \n Скорость ветра (в м/с) - %.1f\n Скорость порывов ветра (в м/с) - %.1f\n Направление ветра - %s",
		request.CityName, r.Geo_object.Province.Name, r.Fact.Wind_speed, r.Fact.Wind_gust, wind_dir)
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
