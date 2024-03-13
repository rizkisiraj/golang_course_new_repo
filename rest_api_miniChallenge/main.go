package main

import (
	"rest_api_mini_challenge/database"
	"rest_api_mini_challenge/routers"
)

func main() {
	database.StartDB()
	var PORT = ":9090"

	r := routers.SetupRouter()

	r.Run(PORT)
}
