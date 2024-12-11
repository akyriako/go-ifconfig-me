package main

import (
	ifconfigme "github.com/akyriako/go-ifconfig-me"
	"log"
)

func main() {
	client := ifconfigme.NewClient()
	response, err := client.Get()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
