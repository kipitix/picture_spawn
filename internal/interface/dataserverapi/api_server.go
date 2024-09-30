package dataserverapi

import (
	"net/http"
)

type apiServer struct {
	mux *http.ServeMux
}

func NewServer() *apiServer {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/random/", randomPictureHandler)

	server := &apiServer{
		mux: mux,
	}

	return server
}

var _ APIServer = (*apiServer)(nil)

func (s *apiServer) ServerMux() *http.ServeMux {
	return s.mux
}
