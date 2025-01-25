package requests

import "time"

type CreateUrl struct {
	url
}

type url struct {
	Url string `json:"url"`
}

type UrlDto struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt"`
}
