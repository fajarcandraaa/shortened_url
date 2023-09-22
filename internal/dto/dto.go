package dto

import "github.com/fajarcandraaa/shortened_url/internal/presentation"

func ToResponse(status string, data interface{}) presentation.Response {
	res := presentation.Response{
		Status: status,
		Data:   data,
	}

	return res
}

func ArrayToResponse(items []interface{}) presentation.Response {
	resp := presentation.Response{
		Data: items,
	}

	return resp
}

func RequestParamToMeta(sortBy, orderBy string, perPage, page int) presentation.MetaPagination {
	resp := presentation.MetaPagination{
		SortBy:  sortBy,
		OrderBy: orderBy,
		PerPage: perPage,
		Page:    page,
	}

	return resp
}
