package main

import (
	"github.com/DiasOrazbaev/lrn-platform/config"
	"github.com/DiasOrazbaev/lrn-platform/pkg/postgres"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(cfg)

	db, err := postgres.NewDatabase(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(db.Db.Config())
}
