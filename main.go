package main

import (
	"./disco"
	"log"
)

func main() {
	discoGateway := disco.NewBuilder().
		ApiKey("9CN7uDxU7HuCe77157F6Ac3eG1Dg17pw").
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
