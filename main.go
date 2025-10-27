package main

import (
	"fmt"

	"github.com/Dima-Melnik/go-restapi-example-with-gin/config"
)

func main() {
	config.LoadConfig("config/config.yaml")
	fmt.Printf("Password: %s", config.InitDatabaseConn())
}
