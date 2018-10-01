package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bloom42/astroflow-go/log"
	"github.com/go-chi/chi"
)

type viewsData struct {
	Date  time.Time `json:"date" gorm:"date"`
	Views uint64    `json:"views" gorm:"views"`
}

type pagesData struct {
	Page    string `json:"page" gorm:"page"`
	Views   uint64 `json:"views" gorm:"views"`
	Uniques uint64 `json:"uniques" gorm:"uniques"`
}

type referrersData struct {
	Referrer string `json:"referrer" gorm:"referrer"`
	Views    uint64 `json:"views" gorm:"views"`
	Uniques  uint64 `json:"uniques" gorm:"uniques"`
}

func (srv *Server) getViewsData(w http.ResponseWriter, r *http.Request) {
	data := []viewsData{}
	id, err := strconv.ParseUint(chi.URLParam(r, "project_id"), 10, 64)
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 400, "invalid {project_id}")
		return
	}

	query := `SELECT date_trunc('day', timestamp) AS "date", count(*) AS "views"
	FROM analytics_events
	WHERE type='page' AND timestamp > current_date - interval '30' day AND project_id = ?
	GROUP BY 1
	ORDER BY date`

	err = srv.DB.Raw(query, id).Scan(&data).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	srv.resJSON(w, 200, data)
}

func (srv *Server) getPagesData(w http.ResponseWriter, r *http.Request) {
	data := []pagesData{}
	id, err := strconv.ParseUint(chi.URLParam(r, "project_id"), 10, 64)
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 400, "invalid {project_id}")
		return
	}

	query := `SELECT page, COUNT(*) AS "views", COUNT(DISTINCT aid) as "uniques"
	FROM (
		SELECT data->>'path' AS "page", anonymous_id AS "aid"
		FROM analytics_events
		WHERE type='page' AND timestamp > current_date - interval '30' day AND project_id = ?
	) AS sub
	GROUP BY page
	ORDER by views DESC
	LIMIT 15;
	`

	err = srv.DB.Raw(query, id).Scan(&data).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	srv.resJSON(w, 200, data)
}

func (srv *Server) getReferrersData(w http.ResponseWriter, r *http.Request) {
	data := []referrersData{}
	id, err := strconv.ParseUint(chi.URLParam(r, "project_id"), 10, 64)
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 400, "invalid {project_id}")
		return
	}

	query := `SELECT referrer, COUNT(*) AS "views", COUNT(DISTINCT aid) as "uniques"
	FROM (
		SELECT data->>'referrer' AS "referrer", anonymous_id AS "aid"
		FROM analytics_events
		WHERE type='page' AND timestamp > current_date - interval '30' day AND project_id = ?
	) AS sub
	GROUP BY referrer
	ORDER by views DESC
	LIMIT 15;
	`
	err = srv.DB.Raw(query, id).Scan(&data).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	for i, datum := range data {
		if datum.Referrer == "" {
			datum.Referrer = "(none)"
			data[i] = datum
			break
		}
	}

	srv.resJSON(w, 200, data)
}
