package api

import "net/http"

func (a *API) newMux() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
