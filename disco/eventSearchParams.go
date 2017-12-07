package disco

/**
 * All possible Query Params for Discovery Event Search.
 * http://developer.ticketmaster.com/products-and-docs/apis/discovery/v2/
 */

const (
/**
 * A string to search against event’s name. Partial word will not be found.
 * ex: keyword=Mado will not found event with name: Madonna
 */
	KEYWORD       = "keyword"
	ATTRACTION_ID = "attractionId"
	VENUE_ID      = "venueId"

	/**
	 * Zipcode or Postal Code of the venue in which the event is taking place.
	 * This is text-based search, not location-based search.
	 * Use lat/long + radius search for nearby events.
	 */
	POSTAL_CODE = "postalCode"
	LAT_LONG    = "latlong"
	RADIUS      = "radius"

	/**
	 * The radius distance unit. Possible values: miles, km.
	 */
	UNIT      = "unit"
	SOURCE    = "source"
	LOCALE    = "locale"
	MARKET_ID = "marketId"

	/**
	 * “2017-01-01T00:00:00Z”
	 */
	START_DATE_TIME = "startDateTime"
	END_DATE_TIME   = "endDateTime"
	INCLUDE_TDB     = "includeTBD"
	INCLUDE_TEST    = "includeTest"

	/**
	 * The number of events returned in the API response. (Max 500)
	 */
	SIZE = "size"

	/**
	 * The page for paginating through the results.
	 */
	PAGE = "page"

	/**
	 * The search sort criteria. Values: “”, “eventDate,date.desc”,
	 * “eventDate,date.asc”, “name,date.desc”, “name,date.asc”.
	 */
	SORT = "sort"

	/**
	 * Include events going onsale after this date.
	 * “2017-01-01T00:00:00Z”
	 */
	ONSALE_START_DATE_TIME = "onsaleStartDateTime"
	ONSALE_END_DATE_TIME   = "onsaleEndDateTime"
	CITY                   = "city"

	/**
	 * ISO value for the country in which you want to query events in.
	 * Possible values are: ‘US’, ‘CA’, ‘AU’, ‘NZ’, ‘MX’.
	 */
	COUNTRY_CODE = "countryCode"
	STATE_CODE   = "stateCode"

	/**
	 * any classification name - segment - genre - sub-genre
	 */
	CLASSIFICATION_NAME = "classificationName"

	CLASSIFICATION_ID = "classificationId"

	SEGMENT_NAME = "segmentName"

	DMA_ID = "dmaId"
)
