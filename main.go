package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marciusvinicius/facebookcollector/"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Hello)
	router.HandleFunc("/profiles/{profileID}", ProfileShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func ProfileShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileID := vars["profileID"]
	profile := Profile{
		Name:      "Marcius",
		FacebokID: "" + profileID,
	}

	facebookConnector = FacebookConector{}
	fields := [2]string
	fields[0] = "first_name"
	result := FacebookConector.Get(fields, profileID)
	json.NewEncoder(w).Encode(result)
}
