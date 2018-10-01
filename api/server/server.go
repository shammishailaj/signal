package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/bloom42/astroflow-go"
	"github.com/bloom42/astroflow-go/log"
	"github.com/bloom42/signal/api/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type APIResponse struct {
	Error  *string     `json:"error"`
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

type Server struct {
	DB        *gorm.DB
	Tmpl      *template.Template
	Router    *chi.Mux
	JWTSecret string
	Host      string
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
		DB:        db,
		Tmpl:      tmpl,
		Router:    router,
		JWTSecret: os.Getenv("JWT_SECRET"),
		Host:      os.Getenv("HOST"),
	}

	router.Use(astroflow.HTTPHandler(log.With()))
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           3600, // Maximum value not ignored by any of major browsers
	})
	router.Use(c.Handler)

	// tracker route
	router.Get("/js", ret.trackerJSRoute)
	// tracking routes
	router.Get("/pixel", ret.pixelGIFRoute)
	router.Post("/events", ret.eventsRoute)

	// API routes
	router.Post("/api/v1/login", ret.loginRoute)

	// authenticated routes
	router.Route("/api/v1", func(r chi.Router) {
		// authenticated routes
		r.Use(ret.authMiddleware)

		r.Get("/accounts", ret.listAccountsRoute)

		r.Get("/projects", ret.listProjectsRoute)
		r.Post("/projects", ret.createProjectRoute)

		r.Get("/projects/{project_id}", ret.getProjectRoute)
		r.Delete("/projects/{project_id}", ret.deleteProjectRoute)

		r.Get("/projects/{project_id}/views", ret.getViewsData)
		r.Get("/projects/{project_id}/pages", ret.getPagesData)
		r.Get("/projects/{project_id}/referrers", ret.getReferrersData)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		ret.resError(w, 404, "route not found")
	})

	return ret, nil
}

func (srv *Server) Run(port string) error {
	return http.ListenAndServe(":"+port, srv.Router)
}

func (srv *Server) AutoMigrate() error {
	return srv.DB.AutoMigrate(
		&db.Project{},
		&db.AnalyticsEvent{},
		&db.Account{},
	).Error
}

// respondwithJSON write json response format
func (srv *Server) resJSON(w http.ResponseWriter, code int, payload interface{}) {
	res := APIResponse{Data: payload, Status: code}
	jsonPayload, err := json.Marshal(res)
	if err != nil {
		log.Error(err.Error())
		srv.resError(w, 500, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonPayload)
}

func (srv *Server) resError(w http.ResponseWriter, code int, message string) {
	if code == 500 {
		message = "internal error"
	}

	res := APIResponse{Error: &message, Status: code}

	jsonPayload, err := json.Marshal(res)
	if err != nil {
		log.Error(err.Error())
	}

	log.With("res", res).Error("error response")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonPayload)
}
