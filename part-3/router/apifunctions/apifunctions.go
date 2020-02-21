package apifunctions

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-3/db"
)

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
		response(w, http.StatusConflict, nil)
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
		response(w, http.StatusNotFound, nil)
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
	}
	err = dataBase.SaveFile()
	if err != nil {
		response(w, http.StatusConflict, nil)
	}
	jsonString, _ := json.Marshal("Item Added")
	response(w, http.StatusOK, jsonString)
}
