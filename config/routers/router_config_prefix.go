package routers

import "github.com/gorilla/mux"

type PathPrefix struct {
	V1       *mux.Router
	ShortUrl *mux.Router
}

func RouterConfigPrefix(se *Serve) *PathPrefix {
	var (
		api      = se.Router.PathPrefix("/api").Subrouter()
		v1       = api.PathPrefix("/v1").Subrouter()
		shorturl = v1.PathPrefix("/urlshort").Subrouter()
	)

	result := &PathPrefix{
		V1:       v1,
		ShortUrl: shorturl,
	}

	return result
}
