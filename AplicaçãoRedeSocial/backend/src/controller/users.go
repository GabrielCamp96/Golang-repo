package controller

import (
	"api/src/db"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user model.User
	if erro = json.Unmarshal(request, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	_, erro = repository.InsertUser(db, user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := db.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	users, erro := repository.GetUsers(db)
	if erro != nil {
		log.Fatal(erro)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		log.Fatal(erro)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 8)
	if erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	user, erro := repository.GetUserById(db, int(ID))
	if erro != nil {
		log.Fatal(erro)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		log.Fatal(erro)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	var user model.User

	ID, erro := strconv.ParseUint(parametros["id"], 10, 8)
	if erro != nil {
		log.Fatal(erro)
	}

	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	if erro := json.Unmarshal(request, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	_, erro = repository.UpdateUser(db, int(ID), user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 8)
	if erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	_, erro = repository.DeleteUser(db, int(ID))
	if erro != nil {
		log.Fatal(erro)
	}

	w.WriteHeader(http.StatusNoContent)
}
