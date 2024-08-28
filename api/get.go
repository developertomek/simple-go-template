package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/developertomek/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id != "" {
		finalId, err := strconv.Atoi(id)
		if err != nil || finalId > len(data.GetAll()) || finalId < 0 {
			http.Error(w, "invalid id", http.StatusBadRequest)
		} else {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
