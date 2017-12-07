package main

import (
	"log"
	"github.com/patjackson52/ticketmaster-discovery-go/disco"
)

func main() {
	discoGateway := disco.NewBuilder().
		ApiKey("{my key}").
		Logging(true).
		Build()

	params := map[string]string{disco.KEYWORD: "Bruno mars"}
	results, err := discoGateway.SearchEvents(params)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(results)
	log.Println(len(results.EmbeddedEvents.Events))
}
