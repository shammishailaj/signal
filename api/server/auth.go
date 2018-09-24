package server

import (
	"net/http"
)

func (srv *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (srv *Server) loginRoute(w http.ResponseWriter, r *http.Request) {
}
