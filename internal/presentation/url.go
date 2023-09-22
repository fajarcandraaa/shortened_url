package presentation

import "time"

// Request
type (
	ShortenedRequest struct {
		URL string `json:"url"`
	}

	ShortenedDetailRequest struct {
		UrlID string `json:"url_id"`
	}
)

// Response
type (
	ShortenedResponse struct {
		ShortURL string `json:"shortened_url"`
	}

	ShortenedDetailResponse struct {
		ShortenedURL string    `json:"shortened_url"`
		TargetUrl    string    `json:"target_url"`
		Clicks       int       `json:"url_click"`
		Latency      int       `json:"url_latency"`
		ExpireTime   time.Time `json:"expire_time"`
	}

	ShortenedListDetailResponse struct {
		ShortenedURL string    `json:"shortened_url"`
		TargetUrl    string    `json:"target_url"`
		Clicks       int       `json:"url_click"`
		Latency      int       `json:"url_latency"`
		ExpireTime   time.Time `json:"expire_time"`
	}
)
