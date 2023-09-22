package routers

import (
	"github.com/fajarcandraaa/shortened_url/internal/repositories"
	"github.com/fajarcandraaa/shortened_url/internal/service"
)

func (se *Serve) initializeRoutes() {

	p := RouterConfigPrefix(se)            // set grouping prefix
	r := repositories.NewRepository(se.DB) //initiate repository
	s := service.NewService(r)             //initiate service

	//initiate endpoint
	shortUrlRoute(p, s)
}
