package dto

import (
	"time"

	"github.com/fajarcandraaa/shortened_url/internal/entity"
	"github.com/fajarcandraaa/shortened_url/internal/presentation"
	"github.com/google/uuid"
)

func UrlShortenedToPayload(url string) presentation.ShortenedRequest {
	resp := presentation.ShortenedRequest{
		URL: url,
	}

	return resp
}

func UrlShortenedToResponse(url string) presentation.ShortenedResponse {
	resp := presentation.ShortenedResponse{
		ShortURL: url,
	}

	return resp
}

func RequestPayloadToDatabase(p presentation.ShortenedRequest, urlShortened string) entity.Url {
	t := time.Now().AddDate(0, 0, 1)
	resp := entity.Url{
		UrlID:        uuid.NewString(),
		ShortenedURL: urlShortened,
		TargetURL:    p.URL,
		UrlClicks:    0,
		UrlLatency:   0,
		ExpireTime:   t,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	return resp
}

func ArrayUrlToResponse(items []entity.Url) presentation.Response {
	resp := presentation.Response{
		Data: items,
	}

	return resp
}
