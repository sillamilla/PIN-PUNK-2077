package main

import (
	"MiniGame-PinUp/MatrixSequence_Service/internal/handler"
	"MiniGame-PinUp/MatrixSequence_Service/internal/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	srv := service.New()
	hnd := handler.New(srv)

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/getSequence", hnd.GetSequence)
		r.Get("/hack", hnd.CallHack)
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
