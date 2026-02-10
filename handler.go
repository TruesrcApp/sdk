package api

import (
	"encoding/json"
	"net/http"
	"truesrc/core"
	"truesrc/storage"

	"github.com/gorilla/mux"
)

var store = storage.NewMemoryStore()

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func CreateNarrative(w http.ResponseWriter, r *http.Request) {
	var n core.Narrative
	json.NewDecoder(r.Body).Decode(&n)

	n = core.PrepareNarrative(n)
	store.Save(n)

	json.NewEncoder(w).Encode(n)
}

func GetNarrative(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	n, ok := store.Get(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(n)
}
