package server

import (
	"net/http"

	"github.com/astrocorp42/astroflow-go/log"
)

type LoginRes struct {
	Token string
}

func (srv *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (srv *Server) loginRoute(w http.ResponseWriter, r *http.Request) {
	log.Info("lol")
	res := LoginRes{"token"}
	srv.resJSON(w, 200, res)
}
