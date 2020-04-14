package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
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

	token := GenerateToken(w)
	log.Print(token) // Temp for debugging

	// TODO:: Send with JWT
	_, err := http.PostForm("http://localhost:8081/toggleWithJWT", url.Values{"k": {"JHDGFUAYEG23RIUETYWERY3RSDFV23RGUE"}})
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintln(w, "Alarm status was successfully switched.")
	}
}

func toggleWithJWT(w http.ResponseWriter, r *http.Request) {
	if validateKeyInUrl(r) {
		superAlarm.status = !superAlarm.status // toggles from on->off and off->on
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/toggle", toggleAlarm).Methods("GET")
	router.HandleFunc("/toggleWithJWT", toggleWithJWT).Methods("POST")

	//log.Fatal(http.ListenAndServeTLS(addrString, "server.crt", "server.key", router))
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", router))
}
