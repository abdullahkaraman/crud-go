package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/abdullahkaraman/crud-go/pkg/mocks"
	"github.com/abdullahkaraman/crud-go/pkg/models"
	"github.com/gorilla/mux"
)

func ReworkPoke(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	var reworkedPoke models.Pokemon
	json.Unmarshal(body, &reworkedPoke)

	for index, poke := range mocks.Pokemons {
		if poke.Id == id {
			poke.Name = reworkedPoke.Name
			poke.Evolve = reworkedPoke.Evolve
			poke.Type = reworkedPoke.Type

			mocks.Pokemons[index] = poke

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Reworked")
			break
		}
	}

}
