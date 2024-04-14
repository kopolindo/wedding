package models

import "log"

func (g *Guest) Print() {
	log.Printf("First name: %s\n", g.FirstName)
	log.Printf("Last name: %s\n", g.LastName)
	log.Printf("UUID: %s\n", g.UUID)
	log.Printf("Confirmed: %t\n", g.Confirmed)
	log.Printf("Number of guests: %d\n", g.NumberOfPartecipants)
	log.Printf("Notes: %s\n", g.Notes)
}
