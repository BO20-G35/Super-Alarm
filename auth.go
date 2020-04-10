package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const passwordFile = "alarmPassword.txt"
const flagFile = "flag.txt"

func validateKeyInUrl(r *http.Request) bool {

	key := r.URL.Query().Get("k")

	rightKey, _ := ioutil.ReadFile(passwordFile)

	rightKey = []byte(strings.TrimSuffix(string(rightKey), "\n"))

	return key == string(rightKey)
}

func GetFlagString() string {
	flag, _ := ioutil.ReadFile(flagFile)
	return string(flag)
}
