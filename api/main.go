package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/astrocorp42/astroflow-go"
	"github.com/astrocorp42/astroflow-go/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var GIF = []byte{
	71, 73, 70, 56, 57, 97, 1, 0, 1, 0, 128, 0, 0, 0, 0, 0,
	255, 255, 255, 33, 249, 4, 1, 0, 0, 0, 0, 44, 0, 0, 0, 0,
	1, 0, 1, 0, 0, 2, 1, 68, 0, 59,
}

var tmpl *template.Template
var db *gorm.DB

func init() {
	var err error

	checkEnv()

	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Config(
		astroflow.SetFormatter(astroflow.NewConsoleFormatter()),
	)

	analytics, err := ioutil.ReadFile("./analytics.min.js")
	if err != nil {
		log.Fatal(err.Error())
	}

	tmpl, err = template.New("analytics").Parse(string(analytics))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func checkEnv() {
	required := []string{
		"DATABASE_URL",
	}

	for _, r := range required {
		if os.Getenv(r) == "" {
			log.With("var", r).Fatal("Missing environment variable")
		}
	}

	// optional env
	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv("PORT", "9090")
	}
}

func main() {
	defer db.Close()
	port := os.Getenv("PORT")
	r := chi.NewRouter()
	c := cors.Default()
	r.Use(c.Handler)

	r.Get("/js", trackerJSRoute)
	r.Get("/pixel.gif", pixelGIFRoute)
	r.Post("/events", eventsRoute)

	log.With("port", port).Info("listenning...")
	log.Fatalf("%s", http.ListenAndServe(":"+port, r))
}
