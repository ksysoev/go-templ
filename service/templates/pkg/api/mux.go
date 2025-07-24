package api

import (
	"net/http"

	"github.com/ksysoev/service/pkg/api/middleware"
)

// newMux creates and returns a new HTTP ServeMux with the API's routes registered.
func (a *API) newMux() *http.ServeMux {
	mux := http.NewServeMux()

	reqIDmid = middleware.NewReqID()

	mux.HandleFunc("/livez", middleware.Use(a.healthCheck))

	return mux
}
