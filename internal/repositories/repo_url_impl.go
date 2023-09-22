package repositories

import (
	"context"

	"github.com/fajarcandraaa/shortened_url/internal/entity"
	"github.com/fajarcandraaa/shortened_url/internal/presentation"
	"github.com/jinzhu/gorm"
)

type ShortenedUrlRepository struct {
	db *gorm.DB
}

func NewShorenedUrlRepository(db *gorm.DB) *ShortenedUrlRepository {
	return &ShortenedUrlRepository{
		db: db,
	}
}

// InsertUrl implements ShortenedUrlRepositoryContract.
func (r *ShortenedUrlRepository) InsertUrl(ctx context.Context, payload entity.Url) error {
	var (
		query = `
			INSERT into urls (url_id, shortened_url, target_url, url_clicks, url_latency, expire_time) 
			VALUES ($1, $2, $3, $4, $5, $6);
		`
	)

	arg := []interface{}{
		&payload.UrlID,
		&payload.ShortenedURL,
		&payload.TargetURL,
		&payload.UrlClicks,
		&payload.UrlLatency,
		&payload.ExpireTime,
	}

	err := r.db.Exec(query, arg...).Error
	if err != nil {
		return err
	}

	return nil
}

// FindUrl implements ShortenedUrlRepositoryContract.
func (r *ShortenedUrlRepository) FindUrl(ctx context.Context, shorUrl string) (*entity.Url, error) {
	var (
		url entity.Url
	)

	err := r.db.First(&url, "shortened_url = ?", shorUrl).Error
	if err != nil {
		return nil, err
	}
	return &url, nil
}

// UpdateClick implements ShortenedUrlRepositoryContract.
func (r *ShortenedUrlRepository) UpdateClick(ctx context.Context, shortUrl string) error {
	var (
		query = `
			UPDATE urls SET url_clicks = url_clicks + 1 WHERE shortened_url = $1
		`
	)

	err := r.db.Exec(query, shortUrl).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateLatency implements ShortenedUrlRepositoryContract.
func (r *ShortenedUrlRepository) UpdateLatency(ctx context.Context, shortUrl string, latency int) error {
	var (
		query = `
			UPDATE urls SET url_latency = $1 WHERE shortened_url = $2
		`
	)

	err := r.db.Exec(query, latency, shortUrl).Error
	if err != nil {
		return err
	}

	return nil
}

// GetUrl implements ShortenedUrlRepositoryContract.
func (r *ShortenedUrlRepository) GetUrl(ctx context.Context, metapayload presentation.MetaPagination) ([]entity.Url, int64, error) {
	var (
		count int64
		urls = make([]entity.Url, 0)

		offset = (metapayload.Page - 1) * metapayload.PerPage
		order  = metapayload.OrderBy + " " + metapayload.SortBy
		model  = r.db.Model(&entity.Url{})
	)

	if err := model.Order(order).Limit(metapayload.PerPage).Offset(offset).Find(&urls).Error; err != nil {
		return nil, 0, err
	}

	_ = model.Count(&count)

	return urls, count, nil
}
