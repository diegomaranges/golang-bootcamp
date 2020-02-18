package utils

import (
	"encoding/json"
	"net/http"
)

/*Generate a JSON Response with status and response arguments*/
func Response(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
