package main

import (
	"log"
	"wedding/src/api"
)

func main() {
	log.Fatal(api.App.Listen(":8080"))
}
