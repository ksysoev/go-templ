package api

import "net/http"

func (a *API) newMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/livez", a.healthCheck)

	return mux
}
