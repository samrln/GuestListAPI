package routes

import (
	"github.com/gorilla/mux"
	"guestListAPI/internal/handlers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/getAllList", handlers.GetAll).Methods("GET")
	router.HandleFunc("/getGuest/{name}", handlers.GetGuest).Methods("GET")
	router.HandleFunc("/getConfirmed", handlers.GetConfirmed).Methods("GET")
	router.HandleFunc("/getNotConfirmed", handlers.GetNotConfirmed).Methods("GET")
	router.HandleFunc("/addGuest", handlers.AddGuest).Methods("POST")
	router.HandleFunc("/updateConfirmed/{name}", handlers.UpdateConfirmed).Methods("PUT")

	return router
}