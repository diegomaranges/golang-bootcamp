package readapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

/*Item algo*/
type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
}

const url = "http://challenge.getsandbox.com/articles"

/*GetAllElements algo*/
func GetAllElements() ([]Item, error) {
	var result []Item
	response, err := http.Get(url)
	if err != nil || response.StatusCode != 200 {
		return result, errors.New("error tring read challenge API")
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil || response.StatusCode != 200 {
		return result, errors.New("error tring read challenge API")
	}
	defer response.Body.Close()
	json.Unmarshal(data, &result)

	return result, nil
}

/*GetElement algo*/
func GetElement(id string) (Item, error) {
	var result Item
	response, err := http.Get(url + "/" + id)
	if err != nil || response.StatusCode != 200 {
		return result, errors.New("error tring read challenge API")
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil || response.StatusCode != 200 {
		return result, errors.New("error tring read challenge API")
	}
	defer response.Body.Close()
	json.Unmarshal(data, &result)

	return result, nil
}
