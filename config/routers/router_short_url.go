package routers

import (
	"github.com/fajarcandraaa/shortened_url/internal/service"
	"github.com/fajarcandraaa/shortened_url/usecase"
)

func shortUrlRoute(p *PathPrefix, s *service.Service) {
	var (
		shortUrlUseCase = usecase.NewShortenedUrlUseCase(s)
	)

	p.ShortUrl.HandleFunc("", shortUrlUseCase.GenerateShortUrl).Methods("POST")
}
