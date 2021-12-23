package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mesaje desde Metodo Get")
}

func POSTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mesaje desde Metodo POST")
}

func PUTUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mesaje desde Metodo PUT")
}

func DELETEUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mesaje desde Metodo DELETE")
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", GetUsers).Methods("Get")       // RUTA PARA LEER USUARIOS
	r.HandleFunc("/api/users", POSTUsers).Methods("POST")     // ruta para crear usuarios
	r.HandleFunc("/api/users", PUTUsers).Methods("PUT")       // ruta para crear usuarios
	r.HandleFunc("/api/users", DELETEUsers).Methods("DELETE") // ruta para crear usuarios

	server := &http.Server{
		Addr:           "8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("listening...")
	server.ListenAndServe()

}
