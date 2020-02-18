package car

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/utils"
)

/*UnitUserActions User interactions with your car*/
type UnitUserActions interface {
	AddItem() error
	ReturnListOfItems(w http.ResponseWriter, r *http.Request)
	ReturnAItem(w http.ResponseWriter, r *http.Request)
	UpdateAItem(w http.ResponseWriter, r *http.Request)
	DeleteAItem(w http.ResponseWriter, r *http.Request)
}

type itemAPI struct {
	id    string
	title string
	price string
}

/*data for a item*/
type item struct {
	title    string
	price    string
	quantity int
}

/*Car Data about diferents cars*/
type Car struct {
	items map[string]*item
}

/*CreateNewCarInstance create new empty car*/
func CreateNewCarInstance() *Car {
	newCar := &Car{}
	newCar.items = make(map[string]*item)
	return newCar
}

/*AddItem Add new item in to the car*/
func (c *Car) AddItem(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	itemID := params["itemId"]
	if c.items[itemID] == nil {
		response, err := http.Get("http://challenge.getsandbox.com/articles")
		if err != nil {
			utils.Response(w, 404, nil)
			return errors.New("map is not initialized")
		}
		bodyResponse, err := ioutil.ReadAll(response.Body)
		if err != nil {
			utils.Response(w, 404, nil)
			return errors.New("map is not initialized")
		}
		var itemsList []itemAPI
		err = json.Unmarshal(bodyResponse, &itemsList)
		if err != nil {
			utils.Response(w, 404, nil)
			return errors.New("map is not initialized")
		}
		fmt.Printf("%+v", itemsList)
		/*c.items[itemID].price = itemsList[itemID].price
		c.items[itemID].title = itemsList[itemID].title
		c.items[itemID].quantity = 1*/
		return nil
	}
	c.items[itemID].quantity++
	return nil
}

/*ReturnListOfItems Return list with all elements in the car*/
func (c *Car) ReturnListOfItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*ReturnAItem Return a single item from the car, if not exist return a error*/
func (c *Car) ReturnAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*UpdateAItem Change item from the car*/
func (c *Car) UpdateAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

/*DeleteAItem Delete item from the car*/
func (c *Car) DeleteAItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}
