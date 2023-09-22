package service

import "github.com/fajarcandraaa/shortened_url/internal/repositories"

type Service struct {
	ShortenedUrlService ShortenedUrlContract
}

func ServiceShortenedUrl(repo *repositories.Repository) ShortenedUrlContract {
	return NewUrlShortenedUrlService(repo)
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		ShortenedUrlService: ServiceShortenedUrl(repo),
	}
}
