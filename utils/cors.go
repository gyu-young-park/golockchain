package utils

import "net/http"

func WrapperCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Add("Conent-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			h(w, r)
		}
	}
}
