package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/VimeanSam/GatorRenter-District-API-Golang/model"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hit home endpoint /")
}

func GetAllDistricts(w http.ResponseWriter, r *http.Request) {
	content, err := json.Marshal(model.Files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//check query parameter
	lte := r.URL.Query()["lte"]

	if len(lte) >= 1 {
		query, err := strconv.ParseFloat(lte[0], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Something wrong...")
			return
		}
		result := make([]model.District, 0)
		for i := 0; i < len(model.Files); i++ {
			str := model.Files[i].Distance_From_SFSU
			spc := strings.Index(str, " ")
			distance, err := strconv.ParseFloat(str[0:spc], 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("Something wrong...")
				return
			}
			if distance <= query {
				result = append(result, model.Files[i])
			}
		}
		content, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(content)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}

func GetPortion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	portions := make([]model.District, 0)
	if strings.Trim(key, " ") == "" {
		log.Println("Please specify district portion")
		return
	}

	for i := 0; i < len(model.Files); i++ {
		if strings.ToLower(model.Files[i].Portion) == strings.ToLower(strings.Trim(key, " ")) {
			//log.Println(files[i])
			portions = append(portions, model.Files[i])
		}
	}

	//if the endpoint does not match any of the district portion. Throw a 404 not found
	if len(portions) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Portion not found"))
		return
	}

	//check query parameter
	lte := r.URL.Query()["lte"]

	if len(lte) >= 1 {
		query, err := strconv.ParseFloat(lte[0], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Something wrong...")
			return
		}
		results := make([]model.District, 0)
		for i := 0; i < len(portions); i++ {
			str := portions[i].Distance_From_SFSU
			spc := strings.Index(str, " ")
			distance, err := strconv.ParseFloat(str[0:spc], 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("Something wrong...")
				return
			}
			if distance <= query {
				results = append(results, portions[i])
			}
		}
		responses, err := json.Marshal(results)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responses)
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
