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

	photoSearchResp, err := fhpApi.SearchPhotosByTerm("bike", values)
	fmt.Println("errors:", err)
	fmt.Printf("%+v\n", photoSearchResp)
}

func photoSearchByTagExample() {
	fhpApi := NewFhpApi()
	values := url.Values{}

	photoSearchResp, err := fhpApi.SearchPhotosByTag("bike", values)
	fmt.Println("errors:", err)
	fmt.Printf("%+v\n", photoSearchResp)
}

func photoSearchByGeoExample() {

	// CN Tower
	latitude := 43.642566
	longitude := -79.387057
	radius := "1km"

	fhpApi := NewFhpApi()
	values := url.Values{}

	photoSearchResp, err := fhpApi.SearchPhotosByGeo(latitude, longitude, radius, values)
	fmt.Println("errors:", err)
	fmt.Printf("%+v\n", photoSearchResp)
}
