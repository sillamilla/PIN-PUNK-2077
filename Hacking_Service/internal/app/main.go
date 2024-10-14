package main

import (
	"MiniGame-PinUp/Hacking_Service/internal/handler"
	"MiniGame-PinUp/Hacking_Service/internal/models"
	"MiniGame-PinUp/Hacking_Service/internal/repository"
	"MiniGame-PinUp/Hacking_Service/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"net/http"
)

func main() {

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	defer db.Close()

	err := db.Model((*models.HackData)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatalf("Migration err: %v", err)
	}

	repo := repository.New(db)
	srv := service.New(repo)
	hnd := handler.New(srv)

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Post("/hack", hnd.Hack)
		r.Get("/getall", hnd.GetAll)
	})

	if err = http.ListenAndServe(":8081", router); err != nil {
		log.Fatal(err)
	}
}
