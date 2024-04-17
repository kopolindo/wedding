package models

import (
	"log"
	"reflect"
)

func (g *Guest) Print() {
	log.Printf("First name: %s\n", g.FirstName)
	log.Printf("Last name: %s\n", g.LastName)
	log.Printf("UUID: %s\n", g.UUID)
	log.Printf("Confirmed: %t\n", g.Confirmed)
	log.Printf("Notes: %s\n", g.Notes)
}

// Function to convert struct to map
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < val.NumField(); i++ {
		result[typ.Field(i).Name] = val.Field(i).Interface()
	}

	return result
}
