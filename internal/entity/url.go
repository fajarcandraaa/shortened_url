package entity

import "time"

type Url struct {
	UrlID        string    `json:"url_id" gorm:"size:36;not null;unique index;primaryKey"`
	ShortenedURL string    `json:"shortened_url" gorm:"size:255;"`
	TargetURL    string    `json:"target_url" gorm:"size:255;"`
	UrlClicks    int       `json:"url_clicks" gorm:"size:20;"`
	UrlLatency   int       `json:"url_latency"`
	ExpireTime   time.Time `json:"expire_time"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
