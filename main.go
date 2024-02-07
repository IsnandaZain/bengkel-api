package main

import "bengkel-api/routers"

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
