package apifunctions

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db/mongodb"
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
			target:   "http://localhost:8080/cars/test",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/test",
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
		assert.Equal(t, table.status, table.response.Code, req.URL.Path)

		_, err := mongodb.CreateNewDBInstance("test", false)
		assert.NoError(t, err, err /*, req.URL.Path*/)
	}

	db, err := mongodb.CreateNewDBInstance("test", false)
	assert.NoError(t, err, err)
	assert.NoError(t, db.DeleteColection(), "error")
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
			target:   "http://localhost:8080/cars/testR",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/test",
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

		tempBody, err := ioutil.ReadAll(table.response.Body)
		assert.NoError(t, err, "error read body")
		items := &fileinteraction.Items{}

		json.Unmarshal(tempBody, items)
		for i, item := range *items {
			assert.Equal(t, strconv.Itoa(i+1), item.ID, item)
		}
	}
}

func TestDeleteCar(t *testing.T) {
	router := mux.NewRouter()
	response := httptest.NewRecorder()
	router.HandleFunc("/cars/{carID}", CreateNewCar).Methods(http.MethodPost)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/cars/testD", nil)
	router.ServeHTTP(response, req)

	temp, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, err, err)
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
			target:   "http://localhost:8080/cars/testD",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/testD",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
	}

	router.HandleFunc("/cars/{carID}", DeleteCar).Methods(http.MethodDelete)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		temp, err := ioutil.ReadAll(table.response.Body)
		assert.Equal(t, err, err)
		assert.Equal(t, table.status, table.response.Code, string(temp), req.URL.Path)
	}
}

func TestReturnItem(t *testing.T) {
	tables := []struct {
		id       string
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			id:       "1",
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/testR/1",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			id:       "2",
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/testR/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			id:       "2",
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/testR/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			id:       "",
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/test/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
		{
			id:       "",
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/testR/11",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
		{
			id:       "",
			method:   http.MethodGet,
			target:   "http://localhost:8080/cars/testR/20",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/cars/{carID}/{itemID}", ReturnItem).Methods(http.MethodGet)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		tempBody, err := ioutil.ReadAll(table.response.Body)
		assert.Equal(t, table.status, table.response.Code, string(tempBody), req.URL.Path)
		assert.NoError(t, err, "error read body")
		if table.response.Code == 200 {
			item := &fileinteraction.Item{}

			json.Unmarshal(tempBody, item)
			assert.Equal(t, table.id, item.ID, item)
		}
	}
}

func TestAddItem(t *testing.T) {
	router := mux.NewRouter()
	response := httptest.NewRecorder()
	router.HandleFunc("/cars/{carID}", CreateNewCar).Methods(http.MethodPost)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/cars/testA", nil)
	router.ServeHTTP(response, req)

	tables := []struct {
		id       string
		quantity int
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			id:       "1",
			quantity: 1,
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/testA/1",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			id:       "2",
			quantity: 1,
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/testA/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			id:       "2",
			quantity: 2,
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/testA/2",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			id:       "3",
			quantity: 1,
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/testA/3",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			id:       "3",
			quantity: 1,
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/testA/30",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusConflict,
		},
		{
			id:       "3",
			quantity: 1,
			method:   http.MethodPost,
			target:   "http://localhost:8080/cars/test/3",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
	}
	router.HandleFunc("/cars/{carID}/{itemID}", AddItem).Methods(http.MethodPost)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		data, err := ioutil.ReadAll(table.response.Body)
		assert.NoError(t, err, err)
		assert.Equal(t, table.status, table.response.Code, string(data), req.URL.Path)

		if table.status == http.StatusOK {
			session, err := mgo.Dial("mongodb://localhost")
			assert.NoError(t, err, err)

			var item fileinteraction.Item
			err = session.DB("CarApi").C("CartestA").Find(bson.M{"_id": table.id}).One(&item)
			assert.NoError(t, err, err)
			assert.Equal(t, table.id, item.ID, item)
			assert.Equal(t, table.quantity, item.Quantity, item)
		}
	}

	router = mux.NewRouter()
	response = httptest.NewRecorder()
	router.HandleFunc("/cars/{carID}", DeleteCar).Methods(http.MethodDelete)

	req = httptest.NewRequest(http.MethodDelete, "http://localhost:8080/cars/testA", nil)
	router.ServeHTTP(response, req)
}

func TestUpdateItem(t *testing.T) {
	item1 := &fileinteraction.Item{}

	item1.ID = "1"
	item1.Price = ""
	item1.Quantity = 5
	item1.Title = ""

	item2 := &fileinteraction.Item{}
	item2.ID = "2"
	item2.Price = ""
	item2.Quantity = 3
	item2.Title = ""

	tables := []struct {
		id        string
		method    string
		target    string
		reader    *fileinteraction.Item
		response  *httptest.ResponseRecorder
		status    int
		aQuantity int
		nQuantity int
	}{
		{
			id:        "1",
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/testU/1",
			reader:    item1,
			response:  httptest.NewRecorder(),
			status:    http.StatusOK,
			aQuantity: 1,
			nQuantity: 5,
		},
		{
			id:        "2",
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/testU/2",
			reader:    item2,
			response:  httptest.NewRecorder(),
			status:    http.StatusOK,
			aQuantity: 8,
			nQuantity: 3,
		},
		{
			id:        "",
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/test/9",
			reader:    nil,
			response:  httptest.NewRecorder(),
			status:    http.StatusNotFound,
			aQuantity: 8,
			nQuantity: 3,
		},
		{
			id:        "",
			method:    http.MethodPut,
			target:    "http://localhost:8080/cars/testU/9",
			reader:    nil,
			response:  httptest.NewRecorder(),
			status:    http.StatusConflict,
			aQuantity: 8,
			nQuantity: 3,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/cars/{carID}/{itemID}", UpdateItem).Methods(http.MethodPut)
	for _, table := range tables {
		var req *http.Request
		if strings.Compare(table.id, "1") == 0 {
			item1Bytes, er := json.Marshal(item1)
			assert.NoError(t, er, item1)
			r := bytes.NewReader(item1Bytes)
			json.NewDecoder(r)
			req = httptest.NewRequest(table.method, table.target, r)
			req.Header.Set("Accept", "application/json")
		} else if strings.Compare(table.id, "2") == 0 {
			item2Bytes, er := json.Marshal(item2)
			assert.NoError(t, er, item2)
			r := bytes.NewReader(item2Bytes)
			json.NewDecoder(r)
			req = httptest.NewRequest(table.method, table.target, r)
			req.Header.Set("Accept", "application/json")
		} else {
			req = httptest.NewRequest(table.method, table.target, nil)
		}

		router.ServeHTTP(table.response, req)
		assert.Equal(t, table.status, table.response.Code, table.id)

		/*if was added correctly check the new quantity*/
		if table.response.Code == http.StatusOK {
			session, err := mgo.Dial("mongodb://localhost")
			assert.NoError(t, err, err)

			var item fileinteraction.Item
			err = session.DB("CarApi").C("CartestU").Find(bson.M{"_id": table.id}).One(&item)
			assert.NoError(t, err, err)
			assert.Equal(t, table.id, item.ID, item)
			assert.Equal(t, table.nQuantity, item.Quantity, item)
			item.Quantity = table.aQuantity
			itemBytes, er := json.Marshal(item)
			assert.NoError(t, er, item)
			r := bytes.NewReader(itemBytes)
			json.NewDecoder(r)

			req = httptest.NewRequest(table.method, table.target, r)
			req.Header.Set("Accept", "application/json")

			router.ServeHTTP(table.response, req)
			assert.Equal(t, table.status, table.response.Code, table.response.Body.String(), string(req.Header.Get("accept")), req.URL.Path, req.Body)
		}
	}
}

func TestDeleteItem(t *testing.T) {
	router := mux.NewRouter()
	response := httptest.NewRecorder()
	router.HandleFunc("/cars/{carID}/{itemID}", AddItem).Methods(http.MethodPost)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/cars/testDI/1", nil)
	router.ServeHTTP(response, req)

	/*expected values*/
	tables := []struct {
		method   string
		target   string
		reader   io.Reader
		response *httptest.ResponseRecorder
		status   int
	}{
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/testDI/8",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/testDI/1",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusOK,
		},
		{
			method:   http.MethodDelete,
			target:   "http://localhost:8080/cars/test/8",
			reader:   nil,
			response: httptest.NewRecorder(),
			status:   http.StatusNotFound,
		},
	}
	router.HandleFunc("/cars/{carID}/{itemID}", DeleteItem).Methods(http.MethodDelete)

	for _, table := range tables {
		req := httptest.NewRequest(table.method, table.target, table.reader)
		router.ServeHTTP(table.response, req)

		assert.Equal(t, table.status, table.response.Code, "error: unexpected status from the response")

		/*if was added correctly check the new quantity*/
		if table.response.Code == http.StatusOK {
			router.HandleFunc("/cars/{carID}/{itemID}", ReturnItem).Methods(http.MethodGet)
			newResponse := httptest.NewRecorder()
			newReq := httptest.NewRequest(http.MethodGet, table.target, nil)
			router.ServeHTTP(newResponse, newReq)
			assert.Equal(t, http.StatusNotFound, newResponse.Code, newReq.URL.Path)
		}
	}
}
