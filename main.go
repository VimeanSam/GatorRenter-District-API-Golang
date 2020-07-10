package main

import (
	"log"
	"net/http"
	"os"

	"github.com/VimeanSam/GatorRenter-District-API-Golang/controller"
	"github.com/gorilla/mux"
)

func handleRequests() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}

	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", controller.Home)
	myrouter.HandleFunc("/districts", controller.GetAllDistricts).Methods("GET")
	myrouter.HandleFunc("/districts/{id}", controller.GetPortion).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, myrouter))
}

func main() {
	handleRequests()
}
