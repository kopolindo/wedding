package backend

import "log"

const NUMBEROFINVITES = 80

func init() {
	log.Println("initiating db from guests.csvs")
	createGuestList()
}
