package readapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

/*Item is the structure used for read all different items from challenge API*/
type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
}

const url = "http://challenge.getsandbox.com/articles"

/*GetAllElements (all items from the API) Return a error if the API request fail or can read the json response*/
func GetAllElements() ([]Item, error) {
	var result []Item
	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		return result, errors.New("error tring read challenge API")
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil || response.StatusCode != 200 {
		return result, errors.New("error tring read json response")
	}

	if er := json.Unmarshal(data, &result); er != nil {
		response.Body.Close()
		return result, errors.New("error tring parse json response")
	}

	return result, response.Body.Close()
}

/*GetElement (particular item from the API) Return a error if the API request fail or can read the json response*/
func GetElement(id string) (Item, error) {
	var result Item
	response, err := http.Get(url + "/" + id)
	if err != nil || response.StatusCode != 200 {
		return result, errors.New("error tring read challenge API")
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil || response.StatusCode != 200 {
		return result, errors.New("error tring read json response")
	}

	if er := json.Unmarshal(data, &result); er != nil {
		response.Body.Close()
		return result, errors.New("error tring parse json response")
	}

	return result, response.Body.Close()
}
