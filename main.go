package main

import (
	"log"
	"net/http"

	"./controller"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", controller.Home)
	myrouter.HandleFunc("/districts", controller.GetAllDistricts).Methods("GET")
	myrouter.HandleFunc("/districts/{id}", controller.GetPortion).Methods("GET")
	log.Fatal(http.ListenAndServe(":3005", myrouter))
}

func main() {
	handleRequests()
}
