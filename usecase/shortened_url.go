package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/fajarcandraaa/shortened_url/helpers"
	"github.com/fajarcandraaa/shortened_url/internal/dto"
	"github.com/fajarcandraaa/shortened_url/internal/entity"
	"github.com/fajarcandraaa/shortened_url/internal/presentation"
	"github.com/fajarcandraaa/shortened_url/internal/service"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrExpiredTime:
			responder.ErrorWithStatusCode(w, http.StatusGone, fmt.Sprint(err))
			return
		default:
			responder.ErrorWithStatusCode(w, http.StatusNotFound, fmt.Sprint(err))
			return
		}
	}
}

func (u *ShortenedUrlUseCase) UrlShortList(w http.ResponseWriter, r *http.Request) {
	var (
		responder    = helpers.NewHTTPResponse("listUrl")
		ctx          = context.Background()
		param        = r.URL.Query()
		paramSortBy  = param.Get("sortby")
		paramOrderBy = param.Get("orderby")
		paramPerPage = param.Get("perpage")
		paramPage    = param.Get("page")
	)

	paginationParam, err := helpers.SetDefaultPginationParam(paramPage, paramPerPage, paramOrderBy, paramSortBy)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, "value of query parameters has diferent type")
		return
	}

	sortBy := paginationParam.SortBy
	orderBy := paginationParam.OrderBy
	perPage := paginationParam.PerPage
	page, _ := strconv.Atoi(paginationParam.Page)

	payload := dto.RequestParamToMeta(sortBy, orderBy, int(perPage), page)

	urlData, total, err := u.service.ShortenedUrlService.ListUrl(ctx, payload)
	if err != nil {
		responder.ErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	pagination, err := helpers.GetPagination(helpers.PaginationParams{
		Path:        "list.items",
		Page:        strconv.Itoa(page),
		TotalRows:   int32(total),
		PerPage:     int32(perPage),
		OrderBy:     orderBy,
		SortBy:      sortBy,
		CurrentPage: int32(page),
	})
	if err != nil {
		responder.ErrorJSON(w, http.StatusConflict, "error pagination")
		return
	}

	urlData.Status = "OK"
	responder.SuccessWithMeta(w, urlData, pagination, http.StatusOK, "short url list")
	return
}
