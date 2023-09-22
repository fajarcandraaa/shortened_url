package repositories

import (
	"context"

	"github.com/fajarcandraaa/shortened_url/internal/entity"
	"github.com/fajarcandraaa/shortened_url/internal/presentation"
)

type ShortenedUrlRepositoryContract interface {
	InsertUrl(ctx context.Context, payload entity.Url) error
	FindUrl(ctx context.Context, urlId string) (*entity.Url, error)
	GetUrl(ctx context.Context, metapayload presentation.MetaPagination) ([]entity.Url, int64, error)
	UpdateLatency(ctx context.Context, urlId string, latency int) error
}
