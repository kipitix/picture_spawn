package dataserverapi

import (
	"fmt"
	"net/http"
)

func randomPictureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "got path\n")
}
