package disco


type EventSearchResult struct {
	EmbeddedEvents EmbeddedEvents `json:"_embedded"`
	Page           Page           `json:"page"`
}

type Page struct {
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Number        int `json:"number"`
}

type EmbeddedEvents struct {
	Events []Event `json:"events"`
}

type PublicRe struct {
	StartDateTime string `json:"startDateTime"`
	StartTBD      string `json:"startTDB"`
	EndDateTime   string `json:"endDateTime"`
}

type Presales struct {
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
	Name          string `json:"name"`
}

type Sales struct {
	PublicRe PublicRe   `json:"publicRe"`
	Presales []Presales `json:"presales"`
}

type DateInfo struct {
	LocalDate      string `json:"localDate"`
	LocalTime      string `json:"localTime"`
	DateTime       string `json:"dateTime"`
	DateTBD        string `json:"dateTBD"`
	DateTBA        string `json:"dateTBA"`
	NoSpecificTime string `json:"noSpecificTime"`
}

type Status struct {
	Code string `json:"code"`
}

type Image struct {
	Ratio    string `json:"ratio"`
	Url      string `json:"url"`
	Width    int `json:"width"`
	Height   int `json:"height"`
	Fallback bool   `json:"fallback"`
}

type Dates struct {
	StartDateInfo DateInfo `json:"startDateInfo"`
	EndDateInfo   DateInfo `json:"endDateInfo"`
	TimeZone      string   `json:"timezone"`
	Status        Status   `json:"status"`
}

type Promoter struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PriceRange struct {
	Type     string  `json:"type"`
	Currency string  `json:"currency"`
	Min      float32 `json:"min"`
	Max      float32 `json:"max"`
}

type Self struct {
	Href      string `json:"href"`
	Templated bool   `json:"href"`
}

type Event struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Id   string `json:"id"`
	//Test            string           `json:"test"`
	Url             string           `json:"url"`
	Local           string           `json:"locale"`
	Images          []Image          `json:"images"`
	Sales           Sales            `json:"sales"`
	Dates           Dates            `json:"dates"`
	Classifications []Classification `json:"classifications"`
	Promoter        Promoter         `json:"promoter"`
	Info            string           `json:"info"`
	PleaseNore      string           `json:"pleaseNote"`
	PriceRages      []PriceRange     `json:"priceRanges"`
	Source          string           `json:"source"`
	References      References       `json:"references"`
}

type References struct {
	TicketmasterUs string `json:"ticketmaster-us"`
}

type Links struct {
	Self Self `json:"self"`
}

type City struct {
	Name string `json:"name"`
}

type State struct {
	Name      string `json:"name"`
	StateCode string `json:"stateCode"`
}

type Country struct {
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}

type Address struct {
	Line1 string `json:"line1"`
}

type Location struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type Market struct {
	Id string `json:"id"`
}

type Dmas struct {
	Id int `json:"id"`
}

type Twitter struct {
	Handle string `json:"handle"`
}

type Social struct {
	Twitter Twitter `json:"twitter"`
}

type BoxOfficInfo struct {
	PhoneNumberDetails    string `json:"phoneNumberDetail"`
	OpenHoursDetail       string `json:"openHoursDetail"`
	AcceptedPaymentDetail string `json:"acceptedPaymentDetail"`
	WillCallDetail        string `json:"willCallDetail"`
}

type GeneralInfo struct {
	GeneralRule string `json:"generalRule"`
	ChildRule   string `json:"childRule"`
}

type Venue struct {
	Name          string       `json:"name"`
	Type          string       `json:"type"`
	Id            string       `json:"id"`
	Test          bool         `json:"test"`
	Url           string       `json:"url"`
	Locale        string       `json:"locale"`
	PostalCode    string       `json:"postalCode"`
	TimeZone      string       `json:"timeZone"`
	City          City         `json:"city"`
	State         State        `json:"state"`
	Country       Country      `json:"country"`
	Address       Address      `json:"address"`
	Location      Location     `json:"location"`
	Markets       []Market     `json:"markets"`
	Dmas          []Dmas       `json:"dmas"`
	Social        Social       `json:"social"`
	BoxOfficInfo  BoxOfficInfo `json:"boxOfficeInfo"`
	ParkingDetail string       `json:"parkingDetail"`
	GeneralInfo   GeneralInfo  `json:"generalInfo"`
	Links         Links        `json:"_links"`
	Images        []Image      `json:"images"`
}

type Segment struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Genre struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Type struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SubType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SubGenre struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Classification struct {
	Primary  bool     `json:"primary"`
	Segment  Segment  `json:"segment"`
	Genre    Genre    `json:"genre"`
	SubGenre SubGenre `json:"subGenre"`
	Type     Type     `json:"type"`
	SubType  SubType  `json:"subType"`
}

type Attraction struct {
	Name            string           `json:"name"`
	Type            string           `json:"type"`
	Id              string           `json:"id"`
	Test            string           `json:"test"`
	Url             string           `json:"url"`
	Locale          string           `json:"locale"`
	Images          []Image          `json:"images"`
	Classifications []Classification `json:"classifications"`
	Links           Links            `json:"_links"`
}

type Embedded struct {
	Venues      []Venue      `json:"venues"`
	Attractions []Attraction `json:"attractions"`
}
