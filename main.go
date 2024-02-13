package main

import (
	"go-final-mygram/database"
	"go-final-mygram/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8000")
}
