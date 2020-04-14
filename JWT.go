package main

import (
	"github.com/dgrijalva/jwt-go"
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

func GenerateToken(w http.ResponseWriter) *jwt.Token {

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
		return nil
	}

	// Set client Cookies TODO:: is necessary?
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return token
}
