package main

import (
	"fmt"
	"log"
	"net/http" // esto nos permite crear el servidor wep
	"time"
)

//func holaMundo(w http.ResponseWriter, r *http.Request) { // En el paquete HTPP se llama a la funcion HandleFunc
//fmt.Fprintf(w, "<h1>hola mundo</h1")

//}

func Prueba(w http.ResponseWriter, r *http.Request) { // En el paquete HTPP se llama a la funcion HandleFunc
	fmt.Fprintf(w, "<h1>hola mundo desde /prueba</h1>")

}

func Usuario(w http.ResponseWriter, r *http.Request) { // En el paquete HTPP se llama a la funcion HandleFunc
	fmt.Fprintf(w, "<h1>hola Usuario</h1")

}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)

}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // se crea una ruta mediante la funcion HandleFunc del paquete Htpp
	// 	fmt.Fprintf(w, "<h1>hola mundo</h1>")
	// })

	msg := mensaje{
		msg: "hola mundo de nuevo",
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)

	http.ListenAndServe(":8080", mux) // esto escucha en el puerto que le indiquemos, que en este caso es 8080 y luego se le pasa un enrutador, un servermux

	mux.HandleFunc("/usuario", Usuario)

	mux.HandleFunc("/pruba", Prueba)

	// se crea el servidor para que maneje las peticiones del usuario

	http.ListenAndServe(":8080", mux)

	mux.Handle("/hola", msg)

	// creo una estructura SERVER :

	server := &http.Server{ // & indica que queremos crear un puntero de una estructura Server
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second, // tiempo de lectura que debe esperar el servidor
		WriteTimeout:   10 * time.Second, // tiempo de escritura que debe esperar el servidor
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())

}
