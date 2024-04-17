package main

import (
	"log"
	"wedding/api"
)

func main() {
	log.Fatal(api.App.Listen(":8080"))
}
