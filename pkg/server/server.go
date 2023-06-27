package server

import "net/http"

func New() *http.ServeMux {
	mux := http.NewServeMux()
	// todo: add health-checks, etc
	return mux
}
