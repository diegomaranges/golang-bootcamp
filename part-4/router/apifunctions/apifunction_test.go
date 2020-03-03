package apifunctions

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/fileinteraction"
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

	var itemResponse fileinteraction.Items

	router.HandleFunc("/cars/{carID}/{itemID}", AddItem).Methods(http.MethodPost)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		data, _ := ioutil.ReadAll(table.response.Body)
		assert.Equal(t, table.status, table.response.Code, string(data), req.URL.Path)

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
