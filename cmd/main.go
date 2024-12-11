package main

import (
	ifconfigme "github.com/akyriako/go-ifconfig-me"
	"log"
	"net/http"
	"time"
)

func main() {
	client := ifconfigme.NewClient(
		ifconfigme.WithTimeout(350*time.Millisecond),
		ifconfigme.WithTransport(&http.Transport{}),
	)
	response, err := client.Get()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
