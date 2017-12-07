# ticketmaster-discovery-go

Client library for Ticketmaster Discovery APIs in Go

Usage:

Grab the package:

`go get github.com/patjackson52/ticketmaster-discovery-go/disco`

Create a `DiscoveryGateway` instance with the builder, then fetch results with SearchEvents()

```

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
```
