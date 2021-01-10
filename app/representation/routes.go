package representation

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ERROR 404")
}

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/auth/signup", SignUp).Methods("POST")
	router.HandleFunc("/api/user/auth/signin", SignUp).Methods("POST")
	router.HandleFunc("/api/user/auth/validate", ValidateToken).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	return router
}
