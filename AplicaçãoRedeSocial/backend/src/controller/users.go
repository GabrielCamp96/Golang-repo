package controller

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criar usuario"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuarios"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuario por ID"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alterar usuario"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Exclui usuario"))
}
