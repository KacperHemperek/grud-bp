package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"www.github.com/KacperHemperek/grud/responses"
	"www.github.com/KacperHemperek/grud/types"
)

func WriteJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, responses.InternalServerError)
	}
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Get(p string, f HandlerFunc) {
	path := fmt.Sprintf("GET %s", p)
	http.HandleFunc(path, Handler(f))
}

func Post(p string, f HandlerFunc) {
	path := fmt.Sprintf("POST %s", p)
	http.HandleFunc(path, Handler(f))
}

func Put(p string, f HandlerFunc) {
	path := fmt.Sprintf("PUT %s", p)
	http.HandleFunc(path, Handler(f))
}

func Delete(p string, f HandlerFunc) {
	path := fmt.Sprintf("DELETE %s", p)
	http.HandleFunc(path, Handler(f))
}

func Handler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		if err := h(w, r); err != nil {
			fmt.Printf("ERROR [%s] %s: %s (%s)\n", r.Method, r.URL, err.Error(), time.Since(now))
			if err, ok := err.(types.ApiError); ok {
				WriteJson(w, err.Status, err)
				return
			}

			WriteJson(w, http.StatusInternalServerError, responses.InternalServerError)
			return
		}
		fmt.Printf("INFO [%s] %s (%s)\n", r.Method, r.URL, time.Since(now))
	}
}

func notFoundHandler(_ http.ResponseWriter, _ *http.Request) error {
	return responses.NotFoundError
}

func NotFoundRouter() {
	Get("/", notFoundHandler)
	Post("/", notFoundHandler)
	Put("/", notFoundHandler)
	Delete("/", notFoundHandler)
}
