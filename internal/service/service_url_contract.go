package service

import (
	"context"

	"github.com/fajarcandraaa/shortened_url/internal/presentation"
)

type ShortenedUrlContract interface {
	NewShortenedUrl(ctx context.Context, payload presentation.ShortenedRequest) (*presentation.Response, error)
	DetailUrl(ctx context.Context, urlID string) (*presentation.Response, error)
	ListUrl(ctx context.Context, sortBy, orderBy string, perPage, page int) (*presentation.Response, error)
	GetLatency(ctx context.Context, urlId string, latency int) error
}
