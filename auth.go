package main

import (
	"io/ioutil"
)

const flagFile = "flag.txt"

func GetFlagString() string {
	flag, _ := ioutil.ReadFile(flagFile)
	return string(flag)
}
