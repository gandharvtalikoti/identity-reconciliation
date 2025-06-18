package main

import (
	"bitespeed-identity/database"
	"bitespeed-identity/routes"
)

func main() {
	database.Connect()
	r := routes.SetupRoutes()
	r.Run(":8080")
}
