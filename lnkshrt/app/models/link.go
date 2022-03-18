package models

type Link struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type LinkInternal struct {
	Link
	CreatedDate string `json:"created_at"`
}
