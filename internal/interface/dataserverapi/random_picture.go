package dataserverapi

import (
	"context"
	"net/http"
)

func (s *apiServer) randomPictureHandler(w http.ResponseWriter, r *http.Request) {
	s.repo.GetRandom(context.TODO())
}
