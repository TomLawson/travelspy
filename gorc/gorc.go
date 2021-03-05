package gorc

import (
	"io"
	"net/http"
)

// RestCountries contains the endpoint to make Rest Countries queries against.
type RestCountries struct {
	API string
}

// New returns a new Rest Countries API object to make queries against.
func New() *RestCountries {
	return &RestCountries{
		API: "https://restcountries.eu/rest/v2/",
	}
}

// Query builds a query against the Rest Countries API
func (rc *RestCountries) Query(endpoint string) *[]byte {
	queryString := rc.API + endpoint

	resp, err := http.Get(queryString)
	if err != nil {
		panic(err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return &responseBody
}
