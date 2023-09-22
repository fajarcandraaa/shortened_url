package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	ShortUrlRepo ShortenedUrlRepositoryContract
}

func NewShortenedUrl(db *gorm.DB) ShortenedUrlRepositoryContract {
	return NewShorenedUrlRepository(db)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ShortUrlRepo: NewShortenedUrl(db),
	}
}