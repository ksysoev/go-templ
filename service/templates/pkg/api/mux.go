package api

import (
	"net/http"

	"{{ .Values.repo }}/pkg/api/middleware"
)

// newMux creates and returns a new HTTP ServeMux with the API's routes registered.
func (a *API) newMux() *http.ServeMux {
	mux := http.NewServeMux()

	withReqID := middleware.NewReqID()

	mux.Handle("/livez", middleware.Use(a.healthCheck, withReqID))

	return mux
}
