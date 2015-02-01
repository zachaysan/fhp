package fhp

type user struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Fullname      string `json:"fullname"`
	UserpicUrl    string `json:"userpic_url"`
	UpgradeStatus int    `json:"upgrade_status"`
}
