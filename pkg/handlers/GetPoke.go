package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/abdullahkaraman/crud-go/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln(err)
	}

	for _, pokemon := range mocks.Pokemons {
		if pokemon.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pokemon)
			break
		}
	}
}
