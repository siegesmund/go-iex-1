package goiex

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ReferenceData struct to interface with /ref-data endpoints
type ReferenceData struct {
	iex
}

// Symbols struct
type Symbols []struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Type      string `json:"type"`
	IexID     string `json:"iexId"`
	Region    string `json:"region"`
	Currency  string `json:"currency"`
	IsEnabled bool   `json:"isEnabled"`
}

// IEXSymbols struct
type IEXSymbols []struct {
	Symbol    string `json:"symbol"`
	Date      string `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
}

// InternationalSybmols struct
type InternationalSybmols []struct {
	Symbol    string `json:"symbol"`
	Exchange  string `json:"exchange"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Type      string `json:"type"`
	IexID     string `json:"iexId"`
	Region    string `json:"region"`
	Currency  string `json:"currency"`
	IsEnabled bool   `json:"isEnabled"`
}

// InternationalExchanges struct
type InternationalExchanges []struct {
	Exchange       string `json:"exchange"`
	Region         string `json:"region"`
	Description    string `json:"description"`
	Mic            string `json:"mic"`
	ExchangeSuffix string `json:"exchangeSuffix"`
}

// USExchanges struct
type USExchanges []struct {
	Name   string `json:"name"`
	Mic    string `json:"mic"`
	TapeID string `json:"tapeId"`
	OatsID string `json:"oatsId"`
	Type   string `json:"type"`
}

// USHolidaysAndTradingDates struct
type USHolidaysAndTradingDates []struct {
	Date           string `json:"date"`
	SettlementDate string `json:"settlementDate"`
}

// Sectors struct
type Sectors []struct {
	Name string `json:"name"`
}

// Tags struct
type Tags []struct {
	Name string `json:"name"`
}

// MutualFundSymbols struct
type MutualFundSymbols []struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Type      string `json:"type"`
	IexID     string `json:"iexId"`
	Region    string `json:"region"`
	Currency  string `json:"currency"`
	IsEnabled bool   `json:"isEnabled"`
}

// NewReferenceData return new ReferenceData
func NewReferenceData(token, version string, base *url.URL, httpClient *http.Client) *ReferenceData {
	apiurl, err := url.Parse("ref-data/")
	if err != nil {
		panic(err)
	}
	return &ReferenceData{
		iex{
			token:   token,
			version: version,
			url:     base,
			apiurl:  apiurl,
			client:  httpClient,
		},
	}
}

// Token return token string
func (rd *ReferenceData) Token() string {
	return rd.token
}

// Version return version string
func (rd *ReferenceData) Version() string {
	return rd.version
}

// URL return URL base
func (rd *ReferenceData) URL() *url.URL {
	return rd.url
}

// APIURL return APIURL
func (rd *ReferenceData) APIURL() *url.URL {
	return rd.apiurl
}

// Client return HTTP client
func (rd *ReferenceData) Client() *http.Client {
	return rd.client
}

// Symbols GET /ref-data/symbols
func (rd *ReferenceData) Symbols() (s Symbols, err error) {
	err = get(rd, &s, "symbols", nil)
	return
}

// IEXSymbols GET /ref-data/iex/symbols
func (rd *ReferenceData) IEXSymbols() (is IEXSymbols, err error) {
	err = get(rd, &is, "iex/symbols", nil)
	return
}

// RegionSymbols GET /ref-data/region/{region}/symbols
func (rd *ReferenceData) RegionSymbols(region string) (is InternationalSybmols, err error) {
	endpoint := fmt.Sprintf("region/%s/symbols", region)
	err = get(rd, &is, endpoint, nil)
	return
}

// ExchangeSymbols GET /ref-data/exchange/{exchange}/symbogTls
func (rd *ReferenceData) ExchangeSymbols(exchange string) (is InternationalSybmols, err error) {
	endpoint := fmt.Sprintf("exchange/%s/symbols", exchange)
	err = get(rd, &is, endpoint, nil)
	return
}

// InternationalExchanges GET /ref-data/exchanges
func (rd *ReferenceData) InternationalExchanges() (ie InternationalExchanges, err error) {
	err = get(rd, &ie, "exchanges", nil)
	return
}

// USExchanges GET /ref-data/market/us/exchanges
func (rd *ReferenceData) USExchanges() (ue USExchanges, err error) {
	err = get(rd, &ue, "market/us/exchanges", nil)
	return
}

// USHolidaysAndTradingDates GET /ref-data/us/dates/{type}/{direction}/{last?}/{startDate?}
func (rd *ReferenceData) USHolidaysAndTradingDates(dateType, direction string, opt ...interface{}) (u USHolidaysAndTradingDates, err error) {
	endpoint := fmt.Sprintf("us/dates/%s/%s", dateType, direction)
	if len(opt) > 0 {
		last := opt[0].(int)
		endpoint = fmt.Sprintf("%s/%s", endpoint, strconv.Itoa(last))
	}
	if len(opt) > 1 {
		startDate := opt[1].(string)
		endpoint = fmt.Sprintf("%s/%s", endpoint, startDate)
	}
	err = get(rd, &u, endpoint, nil)
	return
}

// Sectors GET /ref-data/sectors
func (rd *ReferenceData) Sectors() (s Sectors, err error) {
	err = get(rd, &s, "sectors", nil)
	return
}

// Tags GET /ref-data/tags
func (rd *ReferenceData) Tags() (t Tags, err error) {
	err = get(rd, &t, "tags", nil)
	return
}

// MutualFundSymbols GET /ref-data/mutual-funds/symbols
func (rd *ReferenceData) MutualFundSymbols() (mfs MutualFundSymbols, err error) {
	err = get(rd, &mfs, "mutual-funds/symbols", nil)
	return
}
