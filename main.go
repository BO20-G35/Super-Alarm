package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type alarm struct {
	status string
}

var superAlarm = alarm{status: "0"}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Super Alarm IP adress")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, superAlarm.status)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
}

func getCookie(w http.ResponseWriter, r *http.Request) {

	// Assign JWT as cookie to client
	GenerateToken(w)

	// Button toggles
	fmt.Fprintln(w, "<a href=\"http://localhost:8081/toggle?k=JHDGFUAYEG23RIUETYWERY3RSDFV23RGUE\">toggle</a>")
}

func toggle(w http.ResponseWriter, r *http.Request) {

	if validateKeyInUrl(r) {

		if superAlarm.status == "0" {
			superAlarm.status = "1"
		} else {
			superAlarm.status = "0"
		}

		fmt.Fprintln(w, "Successfully toggled alarm. Status: "+superAlarm.status)

		// Checks for JWT token in cookie
		err := ReadJwtToken(w, r)
		if err == nil {

			// Checks if the current JWT token is the same as previousCookie, if so prints flag
			cookie, err := r.Cookie("token")
			if err == nil {
				if CompareCookies(cookie) {
					fmt.Fprintln(w, GetFlagString())
				}
				previousCookie = cookie
			}
		} else {
			fmt.Fprintln(w, err.Error())
		}

	} else {
		w.WriteHeader(401)
		fmt.Fprintln(w, "You are not authorized to make this request.")
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/getCookie", getCookie).Methods("GET")
	router.HandleFunc("/toggle", toggle).Methods("GET")

	//log.Fatal(http.ListenAndServeTLS(addrString, "server.crt", "server.key", router))
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", router))
}
