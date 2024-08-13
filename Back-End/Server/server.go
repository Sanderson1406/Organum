package server

import (
	"log"
	"net/http"
	"time"

	handler "organum/Handler"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartServer() {
	r := mux.NewRouter()

	for _, route := range Routes {
		var h http.Handler = route.Handler
		if route.RequiresAuth {
			h = handler.AuthMiddleware(h)
		}
		r.Handle(route.Path, h).Methods(route.Method)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	s := &http.Server{
		Addr:         "localhost:8084",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
