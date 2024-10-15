package main

import (
	"MiniGame-PinUp/MatrixSequence_Service/app/config"
	"MiniGame-PinUp/MatrixSequence_Service/internal/handler"
	"MiniGame-PinUp/MatrixSequence_Service/internal/service"
	"MiniGame-PinUp/MatrixSequence_Service/pkg/hackService"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	hackServiceClient := hackService.NewClient(cfg.HackService.HackServiceAddress, cfg.HackService.Endpoint)
	srv := service.New(hackServiceClient)
	hnd := handler.New(srv)

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/getSequence", hnd.GenerateMatrixData)
		r.Get("/hack", hnd.CallHack)
	})

	if err := http.ListenAndServe(":"+cfg.HTTP.Port, router); err != nil {
		log.Fatal(err)
	}
}
