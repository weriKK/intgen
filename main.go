package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type NextID struct {
	ID uint32 `json:"id"`
}

func nextIDHandler(w http.ResponseWriter, r *http.Request) {

	id := NextID{
		ID: generateID(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/next-id", nextIDHandler)

	log.Println("Listening on 0.0.0.0:8080")
	http.ListenAndServe("0.0.0.0:8080", r)
}
