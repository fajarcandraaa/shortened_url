package routers

import (
	"github.com/fajarcandraaa/shortened_url/internal/service"
	"github.com/fajarcandraaa/shortened_url/usecase"
)

func shortUrlRoute(se *Serve, s *service.Service) {
	var (
		shortUrlUseCase = usecase.NewShortenedUrlUseCase(s)
	)

	se.Router.HandleFunc("/urlshort", shortUrlUseCase.GenerateShortUrl).Methods("POST")
	se.Router.HandleFunc("/{shorturl}", shortUrlUseCase.RedirectUrl).Methods("GET")
}
