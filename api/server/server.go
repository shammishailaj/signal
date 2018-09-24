package server

import (
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/astrocorp42/astroflow-go"
	"github.com/astrocorp42/astroflow-go/log"
	"github.com/astrocorp42/signal/api/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Tmpl   *template.Template
	Router *chi.Mux
}

var GIF = []byte{
	71, 73, 70, 56, 57, 97, 1, 0, 1, 0, 128, 0, 0, 0, 0, 0,
	255, 255, 255, 33, 249, 4, 1, 0, 0, 0, 0, 44, 0, 0, 0, 0,
	1, 0, 1, 0, 0, 2, 1, 68, 0, 59,
}

func New(databaseURL string) (*Server, error) {
	var ret *Server

	analytics, err := ioutil.ReadFile("./analytics.min.js")
	if err != nil {
		return ret, err
	}

	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return ret, err
	}

	tmpl, err := template.New("analytics").Parse(string(analytics))
	if err != nil {
		return ret, err
	}

	router := chi.NewRouter()
	ret = &Server{
		DB:     db,
		Tmpl:   tmpl,
		Router: router,
	}

	router.Use(astroflow.HTTPHandler(log.With()))
	c := cors.Default()
	router.Use(c.Handler)

	router.Get("/js", ret.trackerJSRoute)
	router.Get("/pixel.gif", ret.pixelGIFRoute)
	router.Post("/events", ret.eventsRoute)

	return ret, nil
}

func (srv *Server) Run(port string) error {
	return http.ListenAndServe(":"+port, srv.Router)
}

func (srv *Server) AutoMigrate() error {
	return srv.DB.AutoMigrate(&db.Project{}, &db.AnalyticsEvent{}).Error
}
