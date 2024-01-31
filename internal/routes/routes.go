package routes

import (
	"github.com/gorilla/mux"
	"guestListAPI/internal/handlers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/GetAllList", handlers.GetAll).Methods("GET")
	router.HandleFunc("/getGuest/{name}", handlers.GetGuest).Methods("GET")
	router.HandleFunc("/GetConfirmed", handlers.GetConfirmed).Methods("GET")
	router.HandleFunc("/GetNotConfirmed", handlers.GetNotConfirmed).Methods("GET")
	router.HandleFunc("/AddGuest", handlers.AddGuest).Methods("POST")
	router.HandleFunc("/UpdateConfirmed/{name}", handlers.UpdateConfirmed).Methods("PUT")

	return router
}