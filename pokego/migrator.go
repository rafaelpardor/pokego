package main

import (
	"log"

	"github.com/rafaelpardor/pokego/pkg/database"
	"github.com/rafaelpardor/pokego/pkg/pokemon"
)

func migrate() {
	db := database.Get()
	err := pokemon.MigrateModel(db)
	if err != nil {
		log.Fatalln("Cannot migrate pokemon: ", err)
	}
}
