package structs

import "time"

type UrlShortnerResponse struct {
	FullUrl                string    `json:"full_url"`
	ShortUrl               string    `json:"short_url"`
	DeletedAt 			   time.Time `json:"deleted_at"`
	CreatedAt              time.Time `json:"created_at"`
}
