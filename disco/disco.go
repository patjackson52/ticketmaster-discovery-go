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
	"strings"
)

/**
 * Interface for all Discovery Api endpoints
 */
type DiscoveryGateway interface {
	SearchEvents(params map[string]string) (*EventSearchResult, error)
	GetEventDetails(eventId string) (*Event, error)
	SearchAttractions(params map[string]string) (*AttractionSearchResult, error)
	SearchVenues(params map[string]string) (*VenueSearchResult, error)
	GetTopPicks(tapId string, params map[string]string) (*TopPicksResponse, error)
	GetInventoryStatusDetails(eventIds []string) (*[]InventoryStatus, error)
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

func (d discoveryGateway) doPostRequest(path string, params map[string]string) ([]byte, error) {
	url := fmt.Sprint(d.baseUrl, path)
	reader := strings.NewReader(path)
	req, reqErr := http.NewRequest("POST", url, reader)
	if reqErr != nil {
		log.Println(reqErr.Error())
		return nil, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", d.apiKey)
	if params != nil {
		for k, v := range params {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	resp, err := http.Post(req.URL.String(), "application/json", reader)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	result, readerErr := ioutil.ReadAll(resp.Body)

	if readerErr != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil

}

func (d discoveryGateway) doGetRequest(path string, params map[string]string) ([]byte, error) {
	url := fmt.Sprint(d.baseUrl, path)
	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("apikey", d.apiKey)
	if params != nil {
		for k, v := range params {
			q.Add(k, v)
		}
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

	if d.log  {
		fmt.Println(string(body))
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Status: ", resp.Status)
		log.Println("Body: ", string(body))
		return nil, errors.New(fmt.Sprint("Error completing discovery request. Status: ", resp.StatusCode))
	}

	return body, nil
}

func (e discoveryGateway) SearchEvents(params map[string]string) (*EventSearchResult, error) {
	params["view"] = "internal"
	body, err := e.doGetRequest("/discovery/v2/events/", params)
	if err != nil {
		return nil, err
	}

	var results EventSearchResult = EventSearchResult{}

	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &results, nil
}

func (d discoveryGateway) SearchAttractions(params map[string]string) (*AttractionSearchResult, error) {
	body, err := d.doGetRequest("/discovery/v2/attractions/", params)
	if err != nil {
		return nil, err
	}

	var results AttractionSearchResult = AttractionSearchResult{}

	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &results, nil
}

func (d discoveryGateway) GetTopPicks(tapEventId string, params map[string]string) (*TopPicksResponse, error) {
	path := fmt.Sprint("/top-picks/v1/events/", tapEventId)
	body, err := d.doGetRequest(path, params)
	if err != nil {
		return nil, err
	}

	var results TopPicksResponse = TopPicksResponse{}

	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &results, nil
	return nil, nil
}
func (d discoveryGateway) GetEventDetails(eventId string) (*Event, error) {
	path := fmt.Sprint("/discovery/v2/events/", eventId)
	body, err := d.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var results Event = Event{}

	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &results, nil
}

func (d discoveryGateway) SearchVenues(params map[string]string) (*VenueSearchResult, error) {
	body, err := d.doGetRequest("/discovery/v2/venues/", params)
	if err != nil {
		return nil, err
	}

	var results VenueSearchResult = VenueSearchResult{}

	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &results, nil
}

func (d discoveryGateway) GetInventoryStatusDetails(eventIds []string) (*[]InventoryStatus, error) {
	eventIdParams := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(eventIds)), ","), "[]")

	body, err := d.doPostRequest("/inventory-status/v1/availability", map[string]string{"events": eventIdParams})
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))


	var results []InventoryStatus

	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &results, nil
}
