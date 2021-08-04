package main

import (
	"cookbook/src/config"
	"cookbook/src/database"
	"cookbook/src/route"
)

func main() {
	config.Carregar()
	db, _ := database.DBConnection()
	route.SetupRoutes(db)

}
