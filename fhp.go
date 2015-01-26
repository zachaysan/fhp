package fhp

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/garyburd/go-oauth/oauth"
)

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.500px.com/v1/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.500px.com/v1/oauth/authorize",
	TokenRequestURI:               "https://api.500px.com/v1/oauth/access_token",
}

type Configuration struct {
	ConsumerKey    string
	ConsumerSecret string
}

var configuration Configuration

func init() {
	filename := "config.json"
	file, err := os.Open(filename)
	decoder := json.NewDecoder(file)

	if err != nil {
		fmt.Println("error:", err)
		println("Problem loading", filename, "aborting")
		return
	}

	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		println("Malformed", filename, "aborting")
		return
	}
}

func Photo(consumerKey string) {
	println(configuration.ConsumerKey)
	println(consumerKey)
}
