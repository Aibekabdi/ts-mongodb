package http

import "net/http"

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/delete/entry", nil)
	mux.HandleFunc("/put/entry", nil)
	mux.HandleFunc("/get/list", nil)
	return mux
}
