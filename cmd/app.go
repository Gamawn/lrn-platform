package main

import (
	"github.com/DiasOrazbaev/lrn-platform/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(cfg)
}
