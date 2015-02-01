package fhp

// TODO: Decide on CreatedAt type

type comment struct {
	Id           int    `json:"id"`
	UserId       int    `json:"user_id"`
	ToWhomUserId int    `json:"to_whom_user_id"`
	Body         string `json:"body"`
	CreatedAt    string `json:"created_at"`
	ParentId     int    `json:"parent_id"`

	User user
}
