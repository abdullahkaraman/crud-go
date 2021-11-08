package main

import (
	"log"
	"net/http"

	"github.com/abdullahkaraman/crud-go/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemons", handlers.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/pokemons", handlers.AddPoke).Methods(http.MethodPost)
	router.HandleFunc("/pokemon/{id}", handlers.AddPoke).Methods(http.MethodGet)
	router.HandleFunc("/pokemon/{id}", handlers.ReworkPoke).Methods(http.MethodPut)
	router.HandleFunc("/pokemon/{id}", handlers.DeletePoke).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":8081", router)
}
