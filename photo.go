package fhp

type photo struct {
	Id                int     `json:"id"`
	UserId            int     `json:"user_id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Lens              string  `json:"lens"`
	FocalLength       string  `json:"focal_length"`
	Iso               string  `json:"iso"`
	ShutterSpeed      string  `json:"shutter_speed"`
	Aperture          string  `json:"aperture"`
	TimesViewed       int     `json:"times_viewed"`
	Rating            float64 `json:"raiting"`
	Status            int     `json:"status"`
	CreatedAt         string  `json:"created_at"`
	Category          int     `json:"category"`
	Location          string  `json:"location"`
	Privacy           bool    `json:"privacy"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	TakenAt           string  `json:"taken_at"`
	HiResUploaded     int     `json:"hi_res_uploaded"`
	ForSale           bool    `json:"for_sale"`
	Width             int     `json:"width"`
	Height            int     `json:"height"`
	VotesCount        int     `json:"votes_count"`
	FavoritesCount    int     `json:"favorites_count"`
	CommentsCount     int     `json:"comments_count"`
	Nsfw              bool    `json:"nsfw"`
	SalesCount        int     `json:"sales_count"`
	ForSaleDate       string  `json:"for_sale_date"`
	HighestRating     float64 `json:"highest_rating"`
	HighestRatingDate string  `json:"highest_rating_date"`
	ImageUrl          string  `json:"image_url"`
	StoreDownload     bool    `json:"store_download"`
	StorePrint        bool    `json:"store_print"`
	Voted             bool    `json:"voted"`
	Favorited         bool    `json:"favorited"`
	Purchased         bool    `json:"purchased"`

	User user
}

type PhotoResp struct {
	Photo    photo
	Comments []comment
}

type PhotoSearchResp struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalItems  int `json:"total_items"`

	Photos []photo
}
