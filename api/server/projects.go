package server

import (
	"net/http"
	"time"

	"github.com/astrocorp42/astroflow-go/log"
	"github.com/astrocorp42/signal/api/db"
)

type listProjectsRes struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	ID        uint      `json:"id"`
}

func (srv *Server) listProjectsRoute(w http.ResponseWriter, r *http.Request) {
	projects := []db.Project{}

	err := srv.DB.Find(&projects).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	res := make([]listProjectsRes, len(projects))

	for i, project := range projects {
		proj := listProjectsRes{
			Name:      project.Name,
			CreatedAt: project.CreatedAt.UTC(),
			ID:        project.ID,
		}
		res[i] = proj
	}

	srv.resJSON(w, 200, res)
}
