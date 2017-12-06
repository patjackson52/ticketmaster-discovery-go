package disco

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type DiscoveryGateway interface {
	SearchEvents(params map[string]string) (*EventSearchResult, error)
}

type discoveryGateway struct {
	apiKey      string
	baseUrl     string
	prettyPrint bool
	logLevel    LogLevel
}

type LogLevel int

const (
	HEADERS LogLevel = 1 + iota
	BODY
	FULL
)

type DiscoGatewayBuilder interface {
	ApiKey(string) DiscoGatewayBuilder
	BaseUrl(string) DiscoGatewayBuilder
	LogLevel(LogLevel) DiscoGatewayBuilder
	PrettyPrint(bool) DiscoGatewayBuilder
	Build() DiscoveryGateway
}

type discoveryGatewayBuilder struct {
	apiKey      string
	baseUrl     string
	prettyPrint bool
	logLevel    LogLevel
}

func (b *discoveryGatewayBuilder) ApiKey(apiKey string) DiscoGatewayBuilder {
	b.apiKey = apiKey
	return b
}

func (b *discoveryGatewayBuilder) BaseUrl(baseUrl string) DiscoGatewayBuilder {
	b.baseUrl = baseUrl
	return b
}

func (b *discoveryGatewayBuilder) LogLevel(level LogLevel) DiscoGatewayBuilder {
	b.logLevel = level
	return b
}

func (b *discoveryGatewayBuilder) PrettyPrint(enable bool) DiscoGatewayBuilder {
	b.prettyPrint = enable
	return b
}

func (b *discoveryGatewayBuilder) Build() DiscoveryGateway {
	return &discoveryGateway{
		apiKey:      b.apiKey,
		baseUrl:     b.baseUrl,
		prettyPrint: b.prettyPrint,
		logLevel:    b.logLevel}
}

func NewBuilder() DiscoGatewayBuilder {
	return &discoveryGatewayBuilder{}
}

func makeRequest(path string, result interface{}) {

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

	if e.logLevel == FULL || e.logLevel == BODY {
		fmt.Println(string(body))
	}

	var results EventSearchResult = EventSearchResult{}

	if resp.StatusCode != http.StatusOK {
		log.Println("Status: ", resp.Status)

		log.Println("Status: ", string(body))
	}
	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &results, nil
}

