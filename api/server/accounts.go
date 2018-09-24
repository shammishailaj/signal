package server

import (
	"net/http"

	"github.com/astrocorp42/astroflow-go/log"
	"github.com/astrocorp42/signal/api/db"
)

func (srv *Server) listAccountsRoute(w http.ResponseWriter, r *http.Request) {
	accounts := []db.Account{}

	err := srv.DB.Find(&accounts).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	res := make([]string, len(accounts))

	for i, account := range accounts {
		res[i] = account.Username
	}

	srv.resJSON(w, 200, res)
}

func main() {
}
