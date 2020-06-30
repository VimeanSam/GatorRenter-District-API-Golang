package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	json.NewEncoder(w).Encode(files)
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/getAllDistricts", getAllDistricts)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func main() {
	handleRequests()
}
