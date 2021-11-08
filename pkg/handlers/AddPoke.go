package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/abdullahkaraman/crud-go/pkg/mocks"
	"github.com/abdullahkaraman/crud-go/pkg/models"
)

func AddPoke(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var pokemon models.Pokemon
	json.Unmarshal(body, &pokemon)

	pokemon.Id = rand.Intn(1000)
	mocks.Pokemons = append(mocks.Pokemons, pokemon)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")

	defer r.Body.Close()
}
