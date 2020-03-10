package apifunctions

/*
import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/mongodb/fileinteraction"
)

func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Item Created"))
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateEndpoint).Methods("GET")
	return router
}

func TestCreateNewCar(t *testing.T) {
	os.Remove("cars/db2.json")

	tables := []struct {
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusCreated,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/cars/{carID}", CreateNewCar).Methods(http.MethodPost)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		temp, _ := ioutil.ReadAll(table.response.Body)
		assert.Equal(t, table.status, table.response.Code, string(temp), req.URL.Path)
	}

	os.Remove("cars/db2.json")
}

func TestReturnCar(t *testing.T) {
	tables := []struct {
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/1",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/1500",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusConflict,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/cars/{carID}", ReturnCar).Methods(http.MethodGet)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		temp, _ := ioutil.ReadAll(table.response.Body)
		assert.Equal(t, table.status, table.response.Code, string(temp), req.URL.Path)
	}
}

func TestDeleteCar(t *testing.T) {
	router := mux.NewRouter()
	response := httptest.NewRecorder()
	router.HandleFunc("/cars/{carID}", CreateNewCar).Methods(http.MethodPost)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/cars/12", nil)
	router.ServeHTTP(response, req)

	temp, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.Code, string(temp), req.URL.Path)

	tables := []struct {
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/12",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/1500",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
	}

	router.HandleFunc("/cars/{carID}", DeleteCar).Methods(http.MethodDelete)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		temp, _ := ioutil.ReadAll(table.response.Body)
		assert.Equal(t, table.status, table.response.Code, string(temp), req.URL.Path)
	}
}

func TestReturnItem(t *testing.T) {
	tables := []struct {
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/1/12",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/1/4",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
		{
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/1500/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusConflict,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/cars/{carID}/{itemID}", ReturnItem).Methods(http.MethodGet)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		temp, _ := ioutil.ReadAll(table.response.Body)
		assert.Equal(t, table.status, table.response.Code, string(temp), req.URL.Path)
	}
}

func TestAddItem(t *testing.T) {
	/*read file and save information from the item 1*/ /*
	router := mux.NewRouter()
	response := httptest.NewRecorder()
	router.HandleFunc("/cars/{carID}/{itemID}", ReturnItem).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/cars/1/1", nil)
	router.ServeHTTP(response, req)

	temp, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.Code, string(temp), req.URL.Path)

	var item fileinteraction.Items
	er := json.Unmarshal(temp, &item)
	assert.NoError(t, er, req.URL.Path)
*/
/*expected values*/ /*
	tables := []struct {
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/1/1",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/1/14",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusConflict,
		},
		{
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/1500/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusConflict,
		},
	}
*/
/*run add function with the expected results*/ /*
	var itemResponse fileinteraction.Items

	router.HandleFunc("/cars/{carID}/{itemID}", AddItem).Methods(http.MethodPost)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		data, _ := ioutil.ReadAll(table.response.Body)
		assert.Equal(t, table.status, table.response.Code, string(data), req.URL.Path)
*/
/*if was added correctly check the new quantity*/ /*
		if table.response.Code == http.StatusOK {
			newResponse := httptest.NewRecorder()
			newReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/cars/1/1", nil)
			router.ServeHTTP(newResponse, newReq)

			data, _ := ioutil.ReadAll(newResponse.Body)
			assert.Equal(t, http.StatusOK, newResponse.Code, newReq.URL.Path)

			er := json.Unmarshal(data, &itemResponse)
			assert.NoError(t, er, itemResponse, req.URL.Path)
			assert.Equal(t, item.Quantity+1, itemResponse.Quantity, newReq.URL)
		}
	}
}

func TestUpdateItem(t *testing.T) {
	item9 := &Item{}
	item2 := &Item{}
*/
/*read file and save information from the item 1*/ /*
	router := mux.NewRouter()
	response := httptest.NewRecorder()
	router.HandleFunc("/cars/{carID}/{itemID}", ReturnItem).Methods(http.MethodGet)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/cars/1/2", nil)
	router.ServeHTTP(response, req)

	temp, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.Code, string(temp), req.URL.Path)

	var item fileinteraction.Items
	er := json.Unmarshal(temp, &item)
	assert.NoError(t, er, req.URL.Path)*/

/*save in a response Item*/ /*
	item2.ID = "2"
	item2.Price = item.Price
	item2.Quantity = item.Quantity
	item2.Title = item.Title
	item2Bytes, er := json.Marshal(item2)
	assert.NoError(t, er, item2)*/

/*load information for a new Item*/ /*
	item9.ID = "9"
	item9.Price = item.Price
	item9.Quantity = item.Quantity
	item9.Title = item.Title
	item9Bytes, er := json.Marshal(item9)
	assert.NoError(t, er, item9)
*/
/*expected values*/ /*
	tables := []struct {
		method    string
		target    string
		reader    []byte
		response  *httptest.ResponseRecorder
		status    int
		newTarget string
	}{
		{
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/1/2",
			reader:    item9Bytes,
			response:  httptest.NewRecorder(),
			status:    http.StatusOK,
			newTarget: "http://localhost:8080/cars/1/9",
		},
		{
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/1/14",
			reader:    nil,
			response:  httptest.NewRecorder(),
			status:    http.StatusConflict,
			newTarget: "",
		},
		{
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/1500/9",
			reader:    nil,
			response:  httptest.NewRecorder(),
			status:    http.StatusConflict,
			newTarget: "",
		},
		{
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/1/9",
			reader:    nil,
			response:  httptest.NewRecorder(),
			status:    http.StatusConflict,
			newTarget: "",
		},
		{
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/1/9",
			reader:    item2Bytes,
			response:  httptest.NewRecorder(),
			status:    http.StatusOK,
			newTarget: "http://localhost:8080/cars/1/2",
		},
	}*/

/*run add function with the expected results*/ /*
	router.HandleFunc("/cars/{carID}/{itemID}", UpdateItem).Methods(http.MethodPut)

	for _, table := range tables {
		r := bytes.NewReader(table.reader)
		json.NewDecoder(r)
		req = httptest.NewRequest(table.method, table.target, r)
		req.Header.Set("Accept", "application/json")

		router.ServeHTTP(table.response, req)
		assert.Equal(t, table.status, table.response.Code, table.response.Body.String(), string(req.Header.Get("accept")), req.URL.Path, req.Body)
*/
/*if was added correctly check the new quantity*/ /*
		if table.response.Code == http.StatusOK {
			newResponse := httptest.NewRecorder()
			newReq := httptest.NewRequest(http.MethodGet, table.target, nil)
			router.ServeHTTP(newResponse, newReq)
			assert.Equal(t, 404, newResponse.Code, newResponse.Code)

		}
	}
}

func TestDeleteItem(t *testing.T) {*/
/*read file and save information from the item 1*/ /*
	router := mux.NewRouter()
	response := httptest.NewRecorder()
	router.HandleFunc("/cars/{carID}/{itemID}", AddItem).Methods(http.MethodPost)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/cars/1/8", nil)
	router.ServeHTTP(response, req)*/

/*expected values*/ /*
	tables := []struct {
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/1500/1",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusConflict,
		},
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/1/8",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/1/8",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
	}
	router.HandleFunc("/cars/{carID}/{itemID}", DeleteItem).Methods(http.MethodDelete)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		assert.Equal(t, table.status, table.response.Code, "error: unexpected status from the response")*/

/*if was added correctly check the new quantity*/ /*
		if table.response.Code == http.StatusOK {
			router.HandleFunc("/cars/{carID}/{itemID}", ReturnItem).Methods(http.MethodGet)
			newResponse := httptest.NewRecorder()
			newReq := httptest.NewRequest(http.MethodGet, "http://localhost:8080/cars/1/8", nil)
			router.ServeHTTP(newResponse, newReq)
			assert.Equal(t, http.StatusNotFound, newResponse.Code, newReq.URL.Path)
		}
	}
}
*/
