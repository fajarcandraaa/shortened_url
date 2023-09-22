package repositories

import (
	"context"

	"github.com/fajarcandraaa/shortened_url/internal/entity"
	"github.com/fajarcandraaa/shortened_url/internal/presentation"
)

type ShortenedUrlRepositoryContract interface {
	InsertUrl(ctx context.Context, payload entity.Url) error
	GetUrl(ctx context.Context, metapayload presentation.MetaPagination) ([]entity.Url, int64, error)
	FindUrl(ctx context.Context, shorUrl string) (*entity.Url, error)
	UpdateClick(ctx context.Context, shortUrl string) error
	UpdateLatency(ctx context.Context, shortUrl string, latency int) error
}
