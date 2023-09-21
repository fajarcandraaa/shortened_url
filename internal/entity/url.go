package entity

import "time"

type Url struct {
	UrlID        string    `json:"url_id" gorm:"size:36;not null;unique index;primaryKey"`
	ShortenedURL string    `json:"shortened_url" gorm:"size:255;"`
	TargetURL    string    `json:"target_url" gorm:"size:255;"`
	Clicks       int       `json:"url_click" gorm:"size:20;"`
	ExpiryTime   time.Time `json:"expiry_time"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
