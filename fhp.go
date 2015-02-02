package fhp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ChimeraCoder/tokenbucket"
	"github.com/garyburd/go-oauth/oauth"
)

var BaseUrl string

const (
	HTTP_GET  = iota
	HTTP_POST = iota
)

type FhpApi struct {
	Credentials          *oauth.Credentials
	queryQueue           chan query
	bucket               *tokenbucket.Bucket
	returnRateLimitError bool
	HttpClient           *http.Client
}

type query struct {
	url         string
	form        url.Values
	data        interface{}
	method      int
	response_ch chan response
}

type response struct {
	data interface{}
	err  error
}

func init() {
	callbackUrl = "http://verify-oauth.herokuapp.com/"
	BaseUrl = "https://api.500px.com"
	filename := "config.json"
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("error:", err)
		println("Problem loading", filename, "aborting")
		return
	}

	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
		println("Malformed", filename, "aborting")
		return
	}

	oauthClient.Credentials.Token = configuration.ConsumerKey
	oauthClient.Credentials.Secret = configuration.ConsumerSecret
}

func (fhpApi *FhpApi) throttledQuery() {
	for query := range fhpApi.queryQueue {
		response_ch := query.response_ch
		err := fhpApi.execQuery(query.url,
			query.form,
			query.data,
			query.method)

		response_ch <- response{query.data, err}
	}
}

func (fhpApi *FhpApi) execQuery(url string,
	form url.Values,
	data interface{},
	method int) error {
	switch method {
	case HTTP_GET:
		return fhpApi.Get(url, form, data)
	case HTTP_POST:
		return fhpApi.Post(url, form, data)
	default:
		return fmt.Errorf("HTTP method not supported")
	}
}

func (fhpApi FhpApi) Get(url string,
	form url.Values,
	data interface{}) error {

	resp, err := oauthClient.Get(fhpApi.HttpClient,
		fhpApi.Credentials, url, form)

	if err != nil {
		return err
	}

	return decodeResponse(resp, data)
}

func (fhpApi FhpApi) Post(url string,
	form url.Values,
	data interface{}) error {
	return fmt.Errorf("HTTP method not supported")
}

func NewFhpApi() *FhpApi {

	queue := make(chan query)

	fhpApi := &FhpApi{&oauth.Credentials{Token: configuration.FinalToken,
		Secret: configuration.FinalSecret}, queue, nil, false, http.DefaultClient}

	go fhpApi.throttledQuery()

	return fhpApi
}

func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		fmt.Println("Status code not 200", resp.StatusCode)
		return fmt.Errorf("Status error")
	}
	return json.NewDecoder(resp.Body).Decode(data)
}

func Run() {
	photoSearchExample()
}
