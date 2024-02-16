package main

import (
	"os"

	"github.com/e-hastono/mygram/database"
	"github.com/e-hastono/mygram/routers"
)

func main() {
	database.StartDB()

	routerPort := envPortOr("3000")

	routers.StartServer().Run(routerPort)
}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
