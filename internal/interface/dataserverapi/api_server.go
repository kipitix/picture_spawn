package dataserverapi

import (
	"net/http"

	"github.com/kipitix/picture_spawn/internal/domain/imginfo"
)

type apiServer struct {
	mux  *http.ServeMux
	repo imginfo.ImageRepo
}

func NewServer(repo imginfo.ImageRepo) *apiServer {
	mux := http.NewServeMux()

	srv := &apiServer{
		mux:  mux,
		repo: repo,
	}

	// mux.HandleFunc("GET /v1/random/", srv.randomPictureHandler)

	return srv
}

var _ APIServer = (*apiServer)(nil)

func (s *apiServer) ServerMux() *http.ServeMux {
	return s.mux
}
