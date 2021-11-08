package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/abdullahkaraman/crud-go/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeletePoke(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln(err)
	}

	for index, pokemon := range mocks.Pokemons {
		if pokemon.Id == id {
			mocks.Pokemons = append(mocks.Pokemons[:index], mocks.Pokemons[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
