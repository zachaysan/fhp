package fhp

type photo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	User     user
}

type PhotoResp struct {
	Photo photo
}