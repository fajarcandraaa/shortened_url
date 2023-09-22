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
func (r *ShortenedUrlRepository) FindUrl(ctx context.Context, urlId string) (*entity.Url, error) {
	panic("unimplemented")
}

// GetUrl implements ShortenedUrlRepositoryContract.
func (r *ShortenedUrlRepository) GetUrl(ctx context.Context, metapayload presentation.MetaPagination) ([]entity.Url, int64, error) {
	panic("unimplemented")
}

// UpdateLatency implements ShortenedUrlRepositoryContract.
func (r *ShortenedUrlRepository) UpdateLatency(ctx context.Context, urlId string, latency int) error {
	panic("unimplemented")
}
