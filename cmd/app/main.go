package main

import (
	"radovid/internal/database"
	"radovid/router"
)

func main() {
	router := router.SetupRouter()

	database.Initialize()

	router.Run(":8080")
}
