package server

import (
	"net/http"
	"time"

	"github.com/astrocorp42/astroflow-go/log"
)

type viewsData struct {
	Date  time.Time `json:"date" gorm:"date"`
	Views uint64    `json:"views" gorm:"views"`
}

type pagesData struct {
	Page  string `json:"page" gorm:"page"`
	Views uint64 `json:"views" gorm:"views"`
}

type referrersData struct {
	Referrer string `json:"referrer" gorm:"referrer"`
	Views    uint64 `json:"views" gorm:"views"`
}

func (srv *Server) getViewsData(w http.ResponseWriter, r *http.Request) {
	data := []viewsData{}

	query := `SELECT date_trunc('day', timestamp) AS "date", count(*) AS "views"
	FROM analytics_events
	WHERE type='page_view' AND timestamp > current_date - interval '30' day
	GROUP BY 1
	ORDER BY date`

	err := srv.DB.Raw(query).Scan(&data).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	srv.resJSON(w, 200, data)
}

func (srv *Server) getPagesData(w http.ResponseWriter, r *http.Request) {
	data := []pagesData{}

	query := `SELECT data->>'path' AS "page", count(*) AS "views"
	FROM analytics_events
	WHERE type='page_view' AND timestamp > current_date - interval '30' day
	GROUP BY page
	ORDER BY views DESC;`

	err := srv.DB.Raw(query).Scan(&data).Error
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, "")
		return
	}

	srv.resJSON(w, 200, data)
}

func (srv *Server) getReferrersData(w http.ResponseWriter, r *http.Request) {
	data := []referrersData{}

	query := `SELECT data->>'referrer' AS "referrer", count(*) AS "views"
	FROM analytics_events
	WHERE type='page_view' AND timestamp > current_date - interval '30' day
	GROUP BY referrer
	ORDER BY views DESC;`

	err := srv.DB.Raw(query).Scan(&data).Error
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
