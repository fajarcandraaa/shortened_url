package service

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/fajarcandraaa/shortened_url/helpers"
	"github.com/fajarcandraaa/shortened_url/internal/dto"
	"github.com/fajarcandraaa/shortened_url/internal/presentation"
	"github.com/fajarcandraaa/shortened_url/internal/repositories"
)

type shortenedUrlService struct {
	repo *repositories.Repository
}

func NewUrlShortenedUrlService(repo *repositories.Repository) *shortenedUrlService {
	return &shortenedUrlService{
		repo: repo,
	}
}

// NewShortenedUrl implements ShortenedUrlContract.
func (s *shortenedUrlService) NewShortenedUrl(ctx context.Context, payload presentation.ShortenedRequest) (*presentation.Response, error) {
	urls := helpers.EnforceHTTP(payload.URL)
	payload.URL = urls
	urlShort := helpers.Base62Encode(rand.Uint64())

	// shortKey := helpers.GenerateShortKey()

	repoPayload := dto.RequestPayloadToDatabase(payload, urlShort)
	err := s.repo.ShortUrlRepo.InsertUrl(ctx, repoPayload)
	if err != nil {
		return nil, err
	}

	short := fmt.Sprintf("http://localhost:8080/%s", urlShort)
	shortenedUrl := dto.UrlShortenedToResponse(short)

	dtoResp := dto.ToResponse("OK", shortenedUrl)
	return &dtoResp, nil

}

// DetailUrl implements ShortenedUrlContract.
func (s *shortenedUrlService) DetailUrl(ctx context.Context, urlID string) (*presentation.Response, error) {
	panic("unimplemented")
}

// GetLatency implements ShortenedUrlContract.
func (s *shortenedUrlService) GetLatency(ctx context.Context, urlId string, latency int) error {
	panic("unimplemented")
}

// ListUrl implements ShortenedUrlContract.
func (s *shortenedUrlService) ListUrl(ctx context.Context, sortBy string, orderBy string, perPage int, page int) (*presentation.Response, error) {
	panic("unimplemented")
}
