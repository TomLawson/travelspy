package gorc

import "encoding/json"

// Country is the definition of a country. It maps to the restcountries.eu API.
type Country struct {
	Name           string     `json:"name"`
	TopLevelDomain []string   `json:"topLevelDomain"`
	Alpha2Code     string     `json:"alpha2Code"`
	Alpha3Code     string     `json:"alpha3Code"`
	CallingCodes   []string   `json:"callingCodes"`
	Capital        string     `json:"capital"`
	AltSpellings   []string   `json:"altSpellings"`
	Region         string     `json:"region"`
	Subregion      string     `json:"subregion"`
	Population     int        `json:"population"`
	LatLong        []float32  `json:"latlng"`
	Denonym        string     `json:"demonym"`
	Area           float32    `json:"area"`
	Gini           float32    `json:"gini"`
	Timezones      []string   `json:"timezones"`
	Borders        []string   `json:"borders"`
	NativeName     string     `json:"nativeName"`
	NumericCode    string     `json:"numericCode"`
	Currencies     []Currency `json:"currencies"`
	Languages      []Language `json:"languages"`
	// Need to decide if this should be in a struct...
	Translations  map[string]string `json:"translations"`
	Flag          string            `json:"flag"`
	RegionalBlocs []RegionalBloc    `json:"regionalBlocs"`
	Cioc          string            `json:"cioc"`
}

// Currency stores the details of a currency.
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Language contains the definition of a language.
type Language struct {
	Iso639_1   string `json:"iso639_1"`
	Iso639_2   string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

// RegionalBloc defines an area that functions as a group, or Bloc.
type RegionalBloc struct {
	Acronym       string   `json:"acronym"`
	Name          string   `json:"name"`
	OtherAcronyms []string `json:"otherAcronyms"`
	OtherNames    []string `json:"otherNames"`
}

// All returns an array of all countries.
func (rc *RestCountries) All() []Country {
	body := rc.Query("all")
	countries := new([]Country)
	err := json.Unmarshal(*body, countries)
	if err != nil {
		panic(err)
	}
	return *countries
}

// Named returns a country matching the specified name
func (rc *RestCountries) Named(name string) Country {
	body := rc.Query("name/" + name + "?fullText=true")
	countries := new([]Country)
	json.Unmarshal(*body, countries)
	return (*countries)[0]
}
