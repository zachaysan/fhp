package fhp

import (
	"fmt"
	"net/url"
)

func photoExample() {
	fhpApi := NewFhpApi()
	values := url.Values{}
	values.Set("comments", "true")

	photoResp, err := fhpApi.GetPhoto(4928401, values)
	fmt.Println("errors:", err)
	fmt.Printf("%+v\n", photoResp)
}

func photoSearchExample() {
	fhpApi := NewFhpApi()
	values := url.Values{}
	values.Set("only", "Black and White")
	values.Set("image_size", "4")

	photoSearchResp, err := fhpApi.SearchPhotosTerm("bike", values)
	fmt.Println("errors:", err)
	fmt.Printf("%+v\n", photoSearchResp)
}
