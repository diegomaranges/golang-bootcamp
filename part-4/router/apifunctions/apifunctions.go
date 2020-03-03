package apifunctions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.corp.globant.com/diego-maranges/GolangBootcamp/part-4/db"
)

const carPath = "carID"
const itemPath = "itemID"
const directoyDB = "cars/"

/*Item is a struct used for read element from the Json requests*/
type Item struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

/*Genere a new error response, encode results how Json*/
func errorResponse(w http.ResponseWriter, status int, results error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	http.Error(w, results.Error(), status)
}

/*Genere a new response, encode results how Json*/
func response(w http.ResponseWriter, status int, results interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func checkIfLogIn(w http.ResponseWriter, r *http.Request) int {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return http.StatusUnauthorized
		}
		// For any other type of error, return a bad request status
		return http.StatusBadRequest
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized
		}
		return http.StatusBadRequest
	}
	if !tkn.Valid {
		return http.StatusUnauthorized
	}

	return http.StatusOK
}

/*CreateNewCar return

status 201 if the car is already exist*/
func CreateNewCar(w http.ResponseWriter, r *http.Request) {
	if _, err := db.CreateNewDBInstance(directoyDB, mux.Vars(r)[carPath], true); err != nil {
		errorResponse(w, http.StatusCreated, err)
		return
	}
	response(w, http.StatusOK, nil)
}

/*ReturnCar return

status 404 if the car does not exist,

status 409 if have any error traing to load the file or read the Map*/
func ReturnCar(w http.ResponseWriter, r *http.Request) {
	dataBase, err := db.CreateNewDBInstance(directoyDB, mux.Vars(r)[carPath], false)
	if err != nil {
		errorResponse(w, http.StatusNotFound, err)
		return
	}

	if er := dataBase.LoadFile(); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	carMap, err := dataBase.ReturnMap()
	if err != nil {
		errorResponse(w, http.StatusConflict, err)
		return
	}
	response(w, http.StatusOK, carMap)
}

/*DeleteCar return

status 404 if the car does not exist or the car does not deleted*/
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	myDB, err := db.CreateNewDBInstance(directoyDB, mux.Vars(r)[carPath], false)
	if err != nil {
		errorResponse(w, http.StatusNotFound, err)
		return
	}

	if er := myDB.DeleteFile(); er != nil {
		errorResponse(w, http.StatusNotFound, er)
		return
	}

	response(w, http.StatusOK, nil)
}

/*ReturnItem return

status 404 if the car does not exist or the car does not have this item

status 409 if have any error load data*/
func ReturnItem(w http.ResponseWriter, r *http.Request) {
	dataBase, err := db.CreateNewDBInstance(directoyDB, mux.Vars(r)[carPath], false)
	if err != nil {
		errorResponse(w, http.StatusNotFound, err)
		return
	}

	if er := dataBase.LoadFile(); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	itemID := mux.Vars(r)[itemPath]
	item, err := dataBase.Retrieve(itemID)
	if err != nil {
		errorResponse(w, http.StatusNotFound, err)
		return
	}
	response(w, http.StatusOK, item)
}

/*AddItem return

status 404 if the car does not exist

status 409 if have any error load/save data, adding the item or create Json response*/
func AddItem(w http.ResponseWriter, r *http.Request) {
	dataBase, err := db.CreateNewDBInstance(directoyDB, mux.Vars(r)[carPath], false)
	if err != nil {
		errorResponse(w, http.StatusNotFound, err)
		return
	}

	if er := dataBase.LoadFile(); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	itemID := mux.Vars(r)[itemPath]
	if er := dataBase.Add(itemID); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	if er := dataBase.SaveFile(); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	jsonString, err := json.Marshal("Item Added")
	if err != nil {
		errorResponse(w, http.StatusConflict, err)
		return
	}
	response(w, http.StatusOK, jsonString)
}

/*UpdateItem return

status 404 if the car does not exist

status 409 if have any error load/save data, updating the item or create/read Json request/response

status 400 if Json received is wrong*/
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	dataBase, err := db.CreateNewDBInstance(directoyDB, mux.Vars(r)[carPath], false)
	if err != nil {
		errorResponse(w, http.StatusNotFound, err)
		return
	}

	if er := dataBase.LoadFile(); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	itemID := mux.Vars(r)[itemPath]
	var item Item
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	if er := json.Unmarshal(data, &item); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	if er := dataBase.Update(itemID, item.ID); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	if er := dataBase.SaveFile(); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	jsonString, err := json.Marshal("Item Added")
	if err != nil {
		errorResponse(w, http.StatusConflict, err)
		return
	}
	response(w, http.StatusOK, jsonString)
}

/*DeleteItem return

status 404 if the car does not exist

status 409 if have any error load/save data, erasing the item or create Json response*/
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	dataBase, err := db.CreateNewDBInstance(directoyDB, mux.Vars(r)[carPath], false)
	if err != nil {
		errorResponse(w, http.StatusNotFound, err)
		return
	}

	if er := dataBase.LoadFile(); er != nil {
		errorResponse(w, http.StatusConflict, er)
		return
	}

	itemID := mux.Vars(r)[itemPath]
	if err := dataBase.Delete(itemID); err != nil {
		errorResponse(w, http.StatusConflict, err)
		return
	}

	if err := dataBase.SaveFile(); err != nil {
		errorResponse(w, http.StatusConflict, err)
		return
	}

	jsonString, err := json.Marshal("Item Added")
	if err != nil {
		errorResponse(w, http.StatusConflict, err)
		return
	}
	response(w, http.StatusOK, jsonString)
}
