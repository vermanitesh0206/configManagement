package routes

import (
	. "my-app/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterConfigRoutes = func(router *mux.Router) {
	router.HandleFunc("/login/", Login).Methods("POST")
	router.HandleFunc("/add/", AddConfig).Methods("POST")
	router.HandleFunc("/search/{configKey}", SearchConfig).Methods("GET")
	router.HandleFunc("/update/", UpdateConfig).Methods("PUT")
	router.HandleFunc("/delete/{configKey}", DeleteConfig).Methods("DELETE")
}
