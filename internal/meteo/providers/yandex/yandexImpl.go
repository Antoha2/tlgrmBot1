package yandex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *yandexImpl) GetWind(apiUrl, apiKey, apiTocken string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Println("http.NewRequest() - ", err)
		return "", err
	}
	req.Header.Set(apiKey, apiTocken)
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

	restResponse := new(Yandex)
	err = json.Unmarshal(body, restResponse)
	if err != nil {
		log.Println("json.Unmarshal() -", err)
		return "", err
	}

	log.Println(restResponse)
	msg := fmt.Sprintf("%v %v", restResponse.Geo_object.Locality.Name, restResponse.Geo_object.Locality.Name)
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
