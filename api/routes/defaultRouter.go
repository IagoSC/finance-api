package routes

import "net/http"

func SetDefaultHandler() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Finance API"))
	}))
	return router
}
