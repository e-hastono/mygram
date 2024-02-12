package main

import (
	"github.com/e-hastono/mygram/database"
	"github.com/e-hastono/mygram/routers"
)

func main() {
	database.StartDB()

	routers.StartServer().Run(":8080")
}
