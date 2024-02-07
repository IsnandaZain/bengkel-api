package main

import (
	"bengkel-api/database"
	"bengkel-api/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
