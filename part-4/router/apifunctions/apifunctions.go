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
		response(w, http.StatusConflict, carMap)
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
	err := dataBase.Add(itemID)
	if err != nil {
		response(w, http.StatusConflict, nil)
		return
	}
	err = dataBase.SaveFile()
	if err != nil {
		response(w, http.StatusConflict, nil)
		return
	}
	jsonString, _ := json.Marshal("Item Added")
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
		response(w, http.StatusBadRequest, data)
		return
	}
	err = json.Unmarshal(data, &item)
	if err != nil {
		response(w, http.StatusConflict, data)
		return
	}

	err = dataBase.Update(itemID, item.ID)
	if err != nil {
		response(w, http.StatusConflict, item.ID)
		return
	}
	err = dataBase.SaveFile()
	if err != nil {
		response(w, http.StatusConflict, nil)
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
	err := dataBase.Delete(itemID)
	if err != nil {
		response(w, http.StatusConflict, nil)
		return
	}
	err = dataBase.SaveFile()
	if err != nil {
		response(w, http.StatusConflict, nil)
		return
	}
	jsonString, _ := json.Marshal("Item Added")
	response(w, http.StatusOK, jsonString)
}
