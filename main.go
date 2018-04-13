package main

import (
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"net/http"
)

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/all", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/all", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/insertUser", insertUser).Methods("POST")
	myRouter.HandleFunc("/insertUser", optionsRequest).Methods("OPTIONS")

	myRouter.HandleFunc("/delete/{id}/", deleteUser).Methods( "DELETE")
	myRouter.HandleFunc("/delete/{id}/", optionsRequest).Methods( "OPTIONS")

	myRouter.HandleFunc("/edit/{id}/", editUser).Methods("POST")
	myRouter.HandleFunc("/edit/{id}/", optionsRequest).Methods("OPTIONS")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Dara API - Mux Routers")
	handleRequests()
}

