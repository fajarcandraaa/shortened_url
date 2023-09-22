package service

import (
	"context"
	"net/http"
	"time"

	"github.com/fajarcandraaa/shortened_url/internal/presentation"
)

type ShortenedUrlContract interface {
	NewShortenedUrl(ctx context.Context, payload presentation.ShortenedRequest) (*presentation.Response, error)
	ShortenedRedirect(ctx context.Context, shortUrl string, startTime time.Time, w http.ResponseWriter, r *http.Request) error
	ListUrl(ctx context.Context, metapayload presentation.MetaPagination) (*presentation.Response, int64, error)
}
