package routes

import "net/http"

func SetDefaultHandler(r *http.ServeMux) {
	r.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Finance API"))
	}))
}
