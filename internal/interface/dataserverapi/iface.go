package dataserverapi

import "net/http"

type APIServer interface {
	ServerMux() *http.ServeMux
}
