package main

import (
	"borsodoy/radovid/internal/database"
	"borsodoy/radovid/router"
)

func main() {
	router := router.SetupRouter()

	database.Initialize()

	router.Run(":8080")
}
