package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/astrocorp42/astroflow-go/log"
	"github.com/astrocorp42/signal/api/db"
)

type projectRes struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	ID        uint      `json:"id"`
}

type createProjectReq struct {
	Name string `json:"name"`
}

func formatProject(project db.Project) projectRes {
	return projectRes{
		Name:      project.Name,
		CreatedAt: project.CreatedAt.UTC(),
		ID:        project.ID,
	}

}

func (srv *Server) listProjectsRoute(w http.ResponseWriter, r *http.Request) {
	projects := []db.Project{}

	err := srv.DB.Find(&projects).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	res := make([]projectRes, len(projects))

	for i, project := range projects {
		res[i] = formatProject(project)
	}

	srv.resJSON(w, 200, res)
}

func (srv *Server) createProjectRoute(w http.ResponseWriter, r *http.Request) {
	var res projectRes
	var req createProjectReq
	var proj db.Project

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 400, "invalid data")
		return
	}

	proj.Name = req.Name

	err = srv.DB.Create(&proj).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	res = formatProject(proj)
	srv.resJSON(w, 201, res)
}
