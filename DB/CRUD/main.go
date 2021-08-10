package main

import (
	"crud/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/usuarios", routes.PostUser).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", routes.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", routes.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", routes.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", routes.DeleteUser).Methods(http.MethodDelete)

	log.Println("Escutando a porta 5000!")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
