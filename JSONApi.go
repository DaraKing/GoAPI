package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
)

func optionsRequest(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusOK)
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

type User struct {
	Id 			string `json:"id"`
	FirstName	string `json:"first_name"`
	LastName  	string `json:"last_name"`
	Age      	int    `json:"age"`
	Email     	string `json:"email"`
	Gender    	string `json:"gender"`
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	setupResponse(&w, r)

	articles, error := QueryItAndJSON()

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error while querying")
	}

	json.NewEncoder(w).Encode(articles)
}

func insertUser(w http.ResponseWriter, r *http.Request)  {

	setupResponse(&w, r)

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	status := insertInto(user.FirstName,user.LastName,user.Age,user.Gender,user.Email)

	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	setupResponse(&w, r)

	vars := mux.Vars(r)
	key := vars["id"]
	id,err := strconv.Atoi(key)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status := deleteUserInDB(id)

	if status{
		w.WriteHeader(http.StatusOK)
	}else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func editUser(w http.ResponseWriter, r *http.Request){
	setupResponse(&w, r)

	vars := mux.Vars(r)
	key := vars["id"]
	id,err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	status := updateUser(id,user.FirstName,user.LastName,user.Age,user.Email,user.Gender)

	if status {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}