package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type alarm struct {
	status bool
}

var superAlarm = alarm{status: false}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Super Alarm IP adress")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, superAlarm.status)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
}

func toggleAlarm(w http.ResponseWriter, r *http.Request) {
	if validateKeyInUrl(r) {
		superAlarm.status = !superAlarm.status // toggles from on->off and off->on
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprintf(w, "Alarm status was successfully switched.")
	} else {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "Invalid Key.")
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/toggle", toggleAlarm).Methods("GET")

	//log.Fatal(http.ListenAndServeTLS(addrString, "server.crt", "server.key", router))
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", router))
}
