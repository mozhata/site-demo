package reply

import (
	"encoding/json"
	"net/http"
)

// Replyer write result to r
type Replyer func(w http.ResponseWriter)

func serverJSON(v interface{}) Replyer {
	return func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(v); err != nil {
			panic(err)
		}
	}
}

func JSON(v interface{}) Replyer {
	return serverJSON(v)
}

func EmptyJSON() Replyer {
	return serverJSON(nil)
}
