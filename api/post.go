package api

import (
	"encoding/json"
	"net/http"

	"github.com/developertomek/museum/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var exibition data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&exibition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.Add(exibition)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
