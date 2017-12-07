/*
Basic Http client for Ticketmaster Discovery Api (http://developer.ticketmaster.com)

Usage:
discoGateway := disco.NewBuilder().
				ApiKey("my key").
				Logging(true).
				Build()
params := map[string]string{disco.KEYWORD: "Bruno mars"}
results, err := discoGateway.SearchEvents(params)

if err != nil {
	log.Println(err.Error())
}

log.Println(results)
*/
package disco

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"github.com/pkg/errors"
)

/**
 * Interface for all Discovery Api endpoints
 */
type DiscoveryGateway interface {
	SearchEvents(params map[string]string) (*EventSearchResult, error)
}

type discoveryGateway struct {
	apiKey  string
	baseUrl string
	log     bool
}

type DiscoGatewayBuilder interface {
	ApiKey(string) DiscoGatewayBuilder
	BaseUrl(string) DiscoGatewayBuilder
	Logging(bool) DiscoGatewayBuilder
	Build() DiscoveryGateway
}

type discoveryGatewayBuilder struct {
	apiKey   string
	baseUrl  string
	logLevel bool
}

func (b *discoveryGatewayBuilder) ApiKey(apiKey string) DiscoGatewayBuilder {
	b.apiKey = apiKey
	return b
}

func (b *discoveryGatewayBuilder) BaseUrl(baseUrl string) DiscoGatewayBuilder {
	b.baseUrl = baseUrl
	return b
}

func (b *discoveryGatewayBuilder) Logging(enabled bool) DiscoGatewayBuilder {
	b.logLevel = enabled
	return b
}

func (b *discoveryGatewayBuilder) Build() DiscoveryGateway {
	return &discoveryGateway{
		apiKey:   b.apiKey,
		baseUrl:  b.baseUrl,
		log: b.logLevel}
}

func NewBuilder() DiscoGatewayBuilder {
	return &discoveryGatewayBuilder{}
}

func (e discoveryGateway) SearchEvents(params map[string]string) (*EventSearchResult, error) {
	req, _ := http.NewRequest("GET", "https://app.ticketmaster.com/discovery/v2/events/", nil)

	q := req.URL.Query()
	q.Add("apikey", e.apiKey)
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	if e.log  {
		fmt.Println(string(body))
	}

	var results EventSearchResult = EventSearchResult{}

	if resp.StatusCode != http.StatusOK {
		log.Println("Status: ", resp.Status)
		log.Println("Body: ", string(body))
		return nil, errors.New(fmt.Sprint("Error completing discovery request. Status: ", resp.StatusCode))
	}
	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &results, nil
}