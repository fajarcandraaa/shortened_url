package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/fajarcandraaa/shortened_url/helpers"
	"github.com/fajarcandraaa/shortened_url/internal/dto"
	"github.com/fajarcandraaa/shortened_url/internal/entity"
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
	urls, err := helpers.EnforceHTTP(payload.URL)
	if err != nil {
		return nil, err
	}
	payload.URL = urls
	urlShort := helpers.Base62Encode(rand.Uint64())

	repoPayload := dto.RequestPayloadToDatabase(payload, urlShort)
	err = s.repo.ShortUrlRepo.InsertUrl(ctx, repoPayload)
	if err != nil {
		return nil, err
	}

	short := fmt.Sprintf("http://localhost:8080/%s", urlShort)
	shortenedUrl := dto.UrlShortenedToResponse(short)

	dtoResp := dto.ToResponse("OK", shortenedUrl)
	return &dtoResp, nil
}

// GetLatency implements ShortenedUrlContract.
func (s *shortenedUrlService) ShortenedRedirect(ctx context.Context, shortUrl string, startTime time.Time, w http.ResponseWriter, r *http.Request) error {
	timenow := time.Now()
	urlDetail, err := s.repo.ShortUrlRepo.FindUrl(ctx, shortUrl)
	if err != nil {
		return err
	}

	exp := timenow.After(urlDetail.ExpireTime)
	if exp {
		return entity.ErrExpiredTime
	}

	err = s.repo.ShortUrlRepo.UpdateClick(ctx, shortUrl)
	if err != nil {
		return err
	}

	http.Redirect(w, r, urlDetail.TargetURL, http.StatusFound)
	go func() {
		latency := time.Since(startTime) / time.Millisecond
		err := s.repo.ShortUrlRepo.UpdateLatency(ctx, shortUrl, int(latency))
		if err != nil {
			log.Fatalf("latency error")
		}
	}()

	return nil
}

// ListUrl implements ShortenedUrlContract.
func (s *shortenedUrlService) ListUrl(ctx context.Context, metapayload presentation.MetaPagination) (*presentation.Response, int64, error) {
	urlList, total, err := s.repo.ShortUrlRepo.GetUrl(ctx, metapayload)
	if err != nil {
		return nil, 0, err
	}

	resp := dto.ArrayUrlToResponse(urlList)
	return &resp, total, nil
}
