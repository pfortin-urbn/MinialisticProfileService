package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")

	//v1 API calls
	router.HandleFunc("/v1/profile", handleProfileNoParams).Methods("GET", "POST", "PUT")
	router.HandleFunc("/v1/profile/{profileId}", handleProfileWithParams).Methods("GET", "DELETE")

	//v2 API calls
	router.HandleFunc("/v2/profile", v2HandleProfileNoParams).Methods("GET")

	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Index Request")
}

func handleProfileNoParams(w http.ResponseWriter, r *http.Request) {
	log.Println("v1 - GET, or POST, or PUT")
}

func handleProfileWithParams(w http.ResponseWriter, r *http.Request) {
	log.Println("v1 - GET/:1, or DELETE/:1")
}

func v2HandleProfileNoParams(w http.ResponseWriter, r *http.Request) {
	log.Println("v2 - GET")
}
