package routes

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"ProfileService/mongo"
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
	fmt.Fprintf(w, "Hello World!!!")
}

func handleProfileNoParams(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProfile(w, r)
	case "POST":
		createProfile(w, r)
	case "PUT":
		updateProfile(w, r)
	}
}

func handleProfileWithParams(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		showProfile(w, r)
	case "DELETE":
		deleteProfile(w, r)
	}
}

//getProfile - Get all profiles from the DB
func getProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("Get All Profiles")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(Profiles)
	json.NewEncoder(w).Encode(mongo.GetProfiles())
}

//createProfile - Create a profile in the DB
func createProfile(w http.ResponseWriter, r *http.Request) {
}

//updateProfile - Update a profile in the DB
func updateProfile(w http.ResponseWriter, r *http.Request) {
}

//showProfile - given the profile ID in the request this returns the profile to the caller
func showProfile(w http.ResponseWriter, r *http.Request) {

}

// deleteProfile - Removes a profile from the DB
func deleteProfile(w http.ResponseWriter, r *http.Request) {
}



func v2HandleProfileNoParams(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		v2GetProfile(w, r)
	}
}


func v2GetProfile(w http.ResponseWriter, r *http.Request) {
}