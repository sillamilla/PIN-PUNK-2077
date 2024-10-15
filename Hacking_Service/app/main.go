package main

import (
	"MiniGame-PinUp/Hacking_Service/app/config"
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
	cfg := config.GetConfig()
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Postgres.Addr,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		Database: cfg.Postgres.DBName,
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
	router.Post("/hack", hnd.Hack)
	router.Get("/getall", hnd.GetAll)

	if err = http.ListenAndServe(":"+cfg.HTTP.Port, router); err != nil {
		log.Fatal(err)
	}
}
