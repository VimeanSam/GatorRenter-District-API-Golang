package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Districts struct {
	Districts []District `json:"districts"`
}

type District struct {
	Name               string `json:"Name"`
	Description        string `json:"Description"`
	Pros               string `json:"Pros"`
	Cons               string `json:"Cons"`
	Portion            string `json:"Portion"`
	Distance_From_SFSU string `json:"Distance_From_SFSU"`
}

func readJson() Districts {
	// Open our jsonFile
	jsonFile, err := os.Open("districts.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened districts.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our districts array
	var districts Districts

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &districts)
	return districts
}

var files = readJson().Districts

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hit home endpoint /")
}

func getAllDistricts(w http.ResponseWriter, r *http.Request) {
	content, err := json.Marshal(files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}

func getPortion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	portions := make([]District, 0)
	if strings.Trim(key, " ") == "" {
		log.Println("Please specify district portion")
		return
	}

	for i := 0; i < len(files); i++ {
		if strings.ToLower(files[i].Portion) == strings.ToLower(strings.Trim(key, " ")) {
			//log.Println(files[i])
			portions = append(portions, files[i])
		}
	}

	//if the endpoint does not match any of the district portion. Throw a 404 not found
	if len(portions) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Portion not found"))
		return
	}

	//write results to json
	content, err := json.Marshal(portions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", home)
	myrouter.HandleFunc("/districts", getAllDistricts).Methods("GET")
	myrouter.HandleFunc("/districts/{id}", getPortion).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", myrouter))
}

func main() {
	handleRequests()
}
