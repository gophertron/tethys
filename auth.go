package main

import (
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	getTemplate(LOGIN_TEMPLATE).Execute(w, nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {

}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("AuthMiddleware")
	next(w, r)
}

func LoggingMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("LoggingMiddleware")
	next(w, r)
}
