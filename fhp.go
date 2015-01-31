package fhp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/ChimeraCoder/tokenbucket"
	"github.com/garyburd/go-oauth/oauth"
)

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.500px.com/v1/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.500px.com/v1/oauth/authorize",
	TokenRequestURI:               "https://api.500px.com/v1/oauth/access_token",
}

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

type Configuration struct {
	ConsumerKey      string
	ConsumerSecret   string
	OAuthToken       string
	OAuthTokenSecret string
	OAuthVerifier    string
	FinalToken       string
	FinalSecret      string
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

var callbackUrl string
var configuration Configuration

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

func get_verify_token() {
	// TODO: Actually implement this for a server context.
	url, secret, err := AuthorizationURL(callbackUrl)
	fmt.Println(url)
	fmt.Println(secret)
	if err != nil {
		fmt.Println(err)
		println("bailing")
		return
	}
	if err == nil {
		println("exiting")
		return
	}
}

func get_final_tokens() {
	// TODO: Actually implement this
	fmt.Println(configuration.OAuthToken)
	fmt.Println(configuration.OAuthTokenSecret)
	fmt.Println(configuration.OAuthVerifier)

	tempCred := &oauth.Credentials{Token: configuration.OAuthToken,
		Secret: configuration.OAuthTokenSecret}

	tokenCred, a, err := oauthClient.RequestToken(http.DefaultClient,
		tempCred,
		configuration.OAuthVerifier)
	fmt.Println(tokenCred, a, err)

}

func (fhpApi *FhpApi) throttledQuery() {
	for query := range fhpApi.queryQueue {
		fmt.Println("throttled query")
		response_ch := query.response_ch
		fmt.Println(query.url)
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

func (fhpApi *FhpApi) GetPhoto(photo_id int) (photoResp PhotoResp, err error) {

	// Replace with actual url values
	values := url.Values{}

	response_ch := make(chan response)

	fhpApi.queryQueue <- query{BaseUrl + "/v1/photos/" + strconv.Itoa(photo_id),
		values,
		&photoResp,
		HTTP_GET,
		response_ch}

	return photoResp, (<-response_ch).err
}

func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		fmt.Println("Status code not 200", resp.StatusCode)
		return fmt.Errorf("Status error")
	}
	return json.NewDecoder(resp.Body).Decode(data)
}

func Run() {
	fhpApi := NewFhpApi()
	photoResp, err := fhpApi.GetPhoto(4928401)
	fmt.Println("finishing Run")
	fmt.Println(err)
	fmt.Println(photoResp)
}

func AuthorizationURL(callback string) (string, *oauth.Credentials, error) {
	tempCred, err := oauthClient.RequestTemporaryCredentials(http.DefaultClient, callback, nil)
	if err != nil {
		return "", nil, err
	}
	return oauthClient.AuthorizationURL(tempCred, nil), tempCred, nil
}
