package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"restapi/repo"
	"strconv"
)

type Respuesta struct {
	Exito   bool   `json:"exito"`
	Mensaje string `json:"mensaje"`
}

func logFatalIf(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func obtenerLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	libros := repo.TodosLibros()
	encoder.Encode(libros)
}

func obtenerLibro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	libro := repo.ObtenerLibro(id)

	encoder.Encode(libro)
}

func crearLibro(w http.ResponseWriter, r *http.Request) {
	var libro repo.Libro
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&libro)

	libro.ID = strconv.Itoa(rand.Intn(100000000))

	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		respuesta := Respuesta{Exito: false, Mensaje: err.Error()}
		encoder.Encode(respuesta)
		log.Fatalln(err)
	}

	repo.InsertarLibro(libro)

	respuesta := Respuesta{
		Exito:   true,
		Mensaje: "Operacion exitosa",
	}

	encoder.Encode(respuesta)
}

func actualizarLibro(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	var libro repo.Libro

	if err := decoder.Decode(&libro); err != nil {
		respuesta := Respuesta{Exito: false, Mensaje: "El formato del libro no es correcto"}
		encoder.Encode(respuesta)
		return
	}

	params := mux.Vars(r)

	id := params["id"]

	if err := repo.ActualizarLibro(id, libro); err != nil {
		respuesta := Respuesta{Exito: false, Mensaje: err.Error()}
		encoder.Encode(respuesta)
		return
	}

	respuesta := Respuesta{
		Exito:   true,
		Mensaje: "Operacion exitosa",
	}

	encoder.Encode(respuesta)
}

func borrarLibro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	w.Header().Set("Content-Type", "application/json")

	err := repo.BorrarLibro(id)
	var respuesta Respuesta

	respuesta.Exito = err == nil

	if err != nil {
		respuesta.Mensaje = err.Error()
	} else {
		respuesta.Mensaje = "Operacion exitosa"
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(respuesta)
}

func tryRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/books/", obtenerLibros).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", obtenerLibro).Methods(http.MethodGet)
	router.HandleFunc("/books/", crearLibro).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", actualizarLibro).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", borrarLibro).Methods(http.MethodDelete)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tryRouter()
}
