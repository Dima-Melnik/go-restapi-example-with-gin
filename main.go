package main

import (
	"log"

	"github.com/Dima-Melnik/go-restapi-example-with-gin/config"
	db "github.com/Dima-Melnik/go-restapi-example-with-gin/internal/database"
	"github.com/Dima-Melnik/go-restapi-example-with-gin/internal/routes"
)

func main() {
	config.LoadConfig("config/config.yaml")

	db.ConnectDB()
	defer db.CloseDB()

	r := routes.Routes()

	port := config.InitServerPort()

	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
