package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/bloom42/astroflow-go/log"
	"github.com/bloom42/signal/api/db"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
)

type projectRes struct {
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ID         uint      `json:"id"`
	TrackingID string    `json:"tracking_id"`
}

type createProjectReq struct {
	Name string `json:"name"`
}

func formatProject(project db.Project) projectRes {
	return projectRes{
		Name:       project.Name,
		CreatedAt:  project.CreatedAt.UTC(),
		ID:         project.ID,
		TrackingID: project.TrackingID,
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

	uid := ksuid.New()

	proj.Name = req.Name
	proj.TrackingID = uid.String()

	err = srv.DB.Create(&proj).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	res = formatProject(proj)
	srv.resJSON(w, 201, res)
}

func (srv *Server) getProjectRoute(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "project_id"), 10, 64)
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 400, "invalid {project_id}")
		return
	}

	project := db.Project{Model: gorm.Model{ID: uint(id)}}
	err = srv.DB.First(&project).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 404, "not found")
		return
	}

	srv.resJSON(w, http.StatusOK, formatProject(project))
}

func (srv *Server) deleteProjectRoute(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "project_id"), 10, 64)
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 400, "invalid {project_id}")
		return
	}

	project := db.Project{Model: gorm.Model{ID: uint(id)}}
	err = srv.DB.Delete(&project).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 404, "not found")
		return
	}

	srv.resJSON(w, http.StatusOK, map[string]string{})
}
