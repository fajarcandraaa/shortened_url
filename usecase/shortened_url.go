package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fajarcandraaa/shortened_url/helpers"
	"github.com/fajarcandraaa/shortened_url/internal/presentation"
	"github.com/fajarcandraaa/shortened_url/internal/service"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
)

type ShortenedUrlUseCase struct {
	service *service.Service
}

func NewShortenedUrlUseCase(service *service.Service) *ShortenedUrlUseCase {
	return &ShortenedUrlUseCase{
		service: service,
	}
}

// GenerateShortUrl is function to generate short url
func (u *ShortenedUrlUseCase) GenerateShortUrl(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("generateShortenedUrl")
		ctx       = context.Background()
		payload   presentation.ShortenedRequest
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.URL, validation.Required),
	)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	resp, err := u.service.ShortenedUrlService.NewShortenedUrl(ctx, payload)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
		return
	}

	responder.SuccessJSON(w, resp, http.StatusCreated, "generate new shortened url success")
	return
}

func (u *ShortenedUrlUseCase) RedirectUrl(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("redirectUrl")
		ctx       = context.Background()
		shortUrl  = mux.Vars(r)["shorturl"]
	)

	startTime := time.Now()

	err := u.service.ShortenedUrlService.ShortenedRedirect(ctx, shortUrl, startTime, w, r)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusNotFound, fmt.Sprint(err))
		return
	}
}
