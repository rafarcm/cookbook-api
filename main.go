package main

import (
	"cookbook/src/config"
	"cookbook/src/database"
	"cookbook/src/route"
	"log"
)

func main() {
	config.Carregar()
	db, _ := database.DBConnection()
	if erro := database.Migrate(db); erro != nil {
		log.Fatal("Erro ao criar as tabelas do sistema", erro)
	}
	route.SetupRoutes(db)
}
