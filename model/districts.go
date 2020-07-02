package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	jsonFile, err := os.Open("file/districts.json")
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

var Files = readJson().Districts
