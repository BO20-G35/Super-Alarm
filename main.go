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

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8081/toggleWithJWT", nil)
	req.RemoteAddr = r.RemoteAddr // TODO:: unsure of this
	cookie := GenerateToken(w)
	req.AddCookie(&cookie)
	_, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
}

func toggleWithJWT(w http.ResponseWriter, r *http.Request) {

	superAlarm.status = !superAlarm.status // toggles from on->off and off->on
	ReadJwtToken(w, r)
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
