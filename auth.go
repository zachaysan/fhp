package fhp

import (
	"fmt"
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
)

var callbackUrl string
var configuration Configuration

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.500px.com/v1/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.500px.com/v1/oauth/authorize",
	TokenRequestURI:               "https://api.500px.com/v1/oauth/access_token",
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

func AuthorizationURL(callback string) (string,
	*oauth.Credentials,
	error) {

	tempCred, err := oauthClient.RequestTemporaryCredentials(http.DefaultClient,
		callback,
		nil)

	if err != nil {
		return "", nil, err
	}

	return oauthClient.AuthorizationURL(tempCred, nil), tempCred, nil
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
