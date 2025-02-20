package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse sends a JSON response to the client
func JSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
