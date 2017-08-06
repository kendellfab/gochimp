package handlers

import (
	"net/http"
)

type CorsHandler struct {
	Origin string
}

func (ch *CorsHandler) CorsMiddleware(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", ch.Origin)
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length")
		fn(w, r)
	}
}
