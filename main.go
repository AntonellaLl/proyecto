package main

import (
	"fmt"
	"log"
	"net/http" // esto nos permite crear el servidor wep
	"time"
)

// func HolaMundo(w http.ResponseWriter, r *http.Request) { // se crea una ruta mediante la funcion HandleFunc del paquete Htpp
// 	fmt.Fprintf(w, "<h1>hola mundo</h1>")
// }

func Prueba(w http.ResponseWriter, r *http.Request) { // se crea una ruta mediante la funcion HandleFunc del paquete Htpp
	fmt.Fprintf(w, "<h1>hola mundo</h1>")
}

func Usuario(w http.ResponseWriter, r *http.Request) { // se crea una ruta mediante la funcion HandleFunc del paquete Htpp
	fmt.Fprintf(w, "<h1>hola mundo</h1>")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)

}

func main() {

	msg := mensaje{
		msg: "hola mundo de nuevo ",
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)

	mux.HandleFunc("/prueba", Prueba)

	mux.HandleFunc("/usuario", Usuario)

	mux.Handle("/hola", msg)

	server := &http.Server{ // con la & se indica que se quiere crear un puntero de una estructura server
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())

}
