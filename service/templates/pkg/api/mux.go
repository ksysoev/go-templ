package api

import "net/http"

// newMux creates and returns a new HTTP ServeMux with the API's routes registered.
func (a *API) newMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/livez", a.healthCheck)

	return mux
}
