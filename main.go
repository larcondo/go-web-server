package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	const port = "3000"
	// Instancia de http.DefaultServerMux
	mux := http.NewServeMux()

	migrate := flag.Bool(
		"migrate", false, "Crea las tablas en la base de datos",
	)

	flag.Parse()
	if *migrate {
		if err := MakeMigrations(); err != nil {
			log.Fatal(err)
		}
	}

	// Ruta a manejar
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/notes", NotesHandler)

	log.Println("Corriendo en http://localhost:" + port)

	// Servidor escuchando en el puerto 8080
	http.ListenAndServe(":"+port, mux)
}

// IndexHandler nos permite manejar la petición a la ruta '/' y retornar "hola mundo" como respuesta al cliente.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("RemoteAddr:", r.RemoteAddr)
	fmt.Println("RequestURI:", r.RequestURI)

	fmt.Fprint(w, "hola mundo! New Server")
}

// GetNotesHandler nos permite manejar las peticiones a la ruta '/notes' con el método GET.
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Puntero a una estructura de tipo Note.
	n := new(Note)

	// Solicitando todas las notas en la base de datos.
	notes, err := n.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Convirtiendo el slice de notas a formato JSON, retorna un []byte y un error.
	j, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Escribiendo el código de respuesta.
	w.WriteHeader(http.StatusOK)

	// Estableciendo el tipo de contenido del cuerpo de la respuesta.
	w.Header().Set("Content-Type", "application/json")

	// Escribiendo la respuesta, es decir nuestro slice de notas en formato JSON.
	w.Write(j)
}

// CreateNotesHandler nos permite manejar las peticiones a la ruta '/notes' con el método POST.
func CreateNotesHandler(w http.ResponseWriter, r *http.Request) {
	var note Note

	// Tomando el cuerpo de la petición, en formato JSON, y decodificándola e la variable note que acabamos de declarar.
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Creamos la nueva nota gracias al método Create.
	err = note.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateNotesHandler nos permite manejar las peticiones a la ruta '/notes' con el método UPDATE.
func UpdateNotesHandler(w http.ResponseWriter, r *http.Request) {
	var note Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Actualizamos la nota correspondiente.
	err = note.Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
	// obtenemos el valor pasado en la url como query correspondiente a id, del tipo /notes?id=3.
	idStr := r.URL.Query().Get("id")

	// Verificamos que no esté vacío.
	if idStr == "" {
		http.Error(w, "Query id es requerido", http.StatusBadRequest)
		return
	}

	// Convertimos el valor obtenido del query a un int, de ser posible.
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Query id debe ser un número", http.StatusBadRequest)
		return
	}

	var note Note

	// Borramos la nota con el id correspondiente.
	err = note.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetNotesHandler(w, r)
	case http.MethodPost:
		CreateNotesHandler(w, r)
	case http.MethodPut:
		UpdateNotesHandler(w, r)
	case http.MethodDelete:
		DeleteNotesHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}
