package apifunctions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db"
)

/*Item algo*/
type Item struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

/*Genere a new response, encode results how Json*/
func response(w http.ResponseWriter, status int, results interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

/*ReturnCar algo*/
func ReturnCar(w http.ResponseWriter, r *http.Request) {
	dataBase := db.CreateNewDBInstance()
	dataBase.LoadFile()

	carMap, err := dataBase.ReturnMap()
	if err != nil {
		response(w, http.StatusConflict, err)
		return
	}
	response(w, http.StatusOK, carMap)
}

/*ReturnItem algo*/
func ReturnItem(w http.ResponseWriter, r *http.Request) {
	dataBase := db.CreateNewDBInstance()
	dataBase.LoadFile()
	itemID := mux.Vars(r)["itemID"]
	item, err := dataBase.Retrieve(itemID)
	if err != nil {
		response(w, http.StatusNotFound, item)
		return
	}
	response(w, http.StatusOK, item)
}

/*AddItem algo*/
func AddItem(w http.ResponseWriter, r *http.Request) {
	dataBase := db.CreateNewDBInstance()
	dataBase.LoadFile()
	itemID := mux.Vars(r)["itemID"]
	if err := dataBase.Add(itemID); err != nil {
		response(w, http.StatusConflict, nil)
		return
	}

	if err := dataBase.SaveFile(); err != nil {
		response(w, http.StatusConflict, nil)
		return
	}

	jsonString, err := json.Marshal("Item Added")
	if err != nil {
		response(w, http.StatusConflict, nil)
		return
	}
	response(w, http.StatusOK, jsonString)
}

/*UpdateItem algo*/
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	dataBase := db.CreateNewDBInstance()
	dataBase.LoadFile()

	itemID := mux.Vars(r)["itemID"]

	var item Item
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response(w, http.StatusBadRequest, err)
		return
	}
	if er := json.Unmarshal(data, &item); er != nil {
		response(w, http.StatusConflict, er)
		return
	}

	if er := dataBase.Update(itemID, item.ID); er != nil {
		response(w, http.StatusConflict, item.ID)
		return
	}

	if er := dataBase.SaveFile(); er != nil {
		response(w, http.StatusConflict, er)
		return
	}
	jsonString, _ := json.Marshal("Item Added")
	response(w, http.StatusOK, jsonString)
}

/*DeleteItem algo*/
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	dataBase := db.CreateNewDBInstance()
	dataBase.LoadFile()
	itemID := mux.Vars(r)["itemID"]

	if err := dataBase.Delete(itemID); err != nil {
		response(w, http.StatusConflict, err)
		return
	}

	if err := dataBase.SaveFile(); err != nil {
		response(w, http.StatusConflict, err)
		return
	}

	jsonString, err := json.Marshal("Item Added")
	if err != nil {
		response(w, http.StatusConflict, err)
		return
	}
	response(w, http.StatusOK, jsonString)
}
