package main

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var jwtKey = []byte("rstydfkrGRGEARFGHAREG")

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type JwtTokenInfo struct {
	UniqueID string `json:"unique_id"`
	jwt.StandardClaims
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateToken(w http.ResponseWriter) http.Cookie {

	expirationTime := time.Now().Add(5 * time.Minute)

	info := &JwtTokenInfo{
		UniqueID: RandString(15),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, info)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	}

	// Set client Cookies
	http.SetCookie(w, &cookie)

	return cookie
}

func ReadJwtToken(w http.ResponseWriter, r *http.Request) {

	// Extract the jwt cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("no cookie set")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		log.Println("cookie parse error")
		return
	}

	// Get the JWT string and parse
	tokenstring := cookie.Value
	jwtInfo := &JwtTokenInfo{}
	tkn, err := jwt.ParseWithClaims(tokenstring, jwtInfo, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("invalid signature")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		log.Println("cannot parse JWT")
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("token is invalid")
		return
	}

	log.Println("Unique ID: " + jwtInfo.UniqueID)

	// TODO:: check if uniqueID from jwt is same as logged cookie
}
