package server

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astrocorp42/astroflow-go/log"
)

type trackerT struct {
	ID string
}

func (srv *Server) trackerJSRoute(w http.ResponseWriter, r *http.Request) {
	trackingID := r.URL.Query().Get("id")
	t := trackerT{trackingID}

	w.Header().Set("Content-Type", "text/javascript")
	err := srv.Tmpl.Execute(w, t)
	if err != nil {
		log.With("tracking_id", trackingID).Error(err.Error())
	}
}

func (srv *Server) pixelGIFRoute(w http.ResponseWriter, r *http.Request) {
	encodedEvents := r.URL.Query().Get("events")
	if encodedEvents != "" {
		decoded, err := base64.StdEncoding.DecodeString(encodedEvents)
		if err != nil {
			log.Error(err.Error())
		} else {
			fmt.Println(decoded)
		}
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Content-Type", "image/gif")
	w.Write(GIF)
}

func (srv *Server) eventsRoute(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println(string(body))
	r.Body.Close()

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}
