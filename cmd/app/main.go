package main

import (
	"borsodoy/radovid/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
