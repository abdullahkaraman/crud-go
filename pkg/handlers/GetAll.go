package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/abdullahkaraman/crud-go/pkg/mocks"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Pokemons)
}
