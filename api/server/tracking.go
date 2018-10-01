package server

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bloom42/astroflow-go/log"
	"github.com/bloom42/signal/api/db"
)

type trackerTmplVar struct {
	ID       string
	Endpoint string
}

func (srv *Server) trackerJSRoute(w http.ResponseWriter, r *http.Request) {
	trackingID := r.URL.Query().Get("id")
	t := trackerTmplVar{ID: trackingID, Endpoint: srv.Host}

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
			parts := strings.Split(r.RemoteAddr, ":")
			if len(parts) == 2 {
				go processEventsPayload(srv, decoded, r.Header.Get("user-agent"), parts[0])
			} else {
				go processEventsPayload(srv, decoded, r.Header.Get("user-agent"), r.RemoteAddr)
			}
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
	r.Body.Close()

	parts := strings.Split(r.RemoteAddr, ":")
	if len(parts) == 2 {
		go processEventsPayload(srv, body, r.Header.Get("user-agent"), parts[0])
	} else {
		go processEventsPayload(srv, body, r.Header.Get("user-agent"), r.RemoteAddr)
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func processEventsPayload(srv *Server, payload []byte, ua, ip string) {
	now := time.Now().UTC()
	events := []Event{}

	err := json.Unmarshal(payload, &events)
	if err != nil {
		log.Error(err.Error())
		return
	}

	if len(events) == 0 {
		return
	}

	var project db.Project
	err = srv.DB.First(&project, "tracking_id = ?", events[0].TrackingID).Error
	if err != nil {
		log.Error(err.Error())
		return
	}

	n := 0
	for _, event := range events {
		if event.Type != "page" {
			continue
		}

		// clean input data
		if s, ok := event.Data["referrer"].(string); ok && s != "" && s != "/" {
			u, err := url.Parse(s)
			if err == nil {
				event.Data["referrer"] = u.Host
			} else {
				event.Data["referrer"] = ""
			}
		}
		if s, ok := event.Data["path"].(string); ok && s != "" && s != "/" {
			event.Data["path"] = strings.TrimRight(s, "/")
		}

		data, err := json.Marshal(event.Data)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		ev := db.AnalyticsEvent{
			ProjectID:   project.ID,
			ReceivedAt:  now,
			UserAgent:   ua,
			IP:          ip,
			TrackingID:  event.TrackingID,
			Type:        event.Type,
			AnonymousID: event.AnonymousID,
			Timestamp:   time.Unix(0, int64(time.Millisecond)*event.Timestamp).UTC(),
			Data:        data,
		}
		err = srv.DB.Create(&ev).Error
		if err != nil {
			log.Error(err.Error())
			continue
		}
		n += 1
	}
	log.With("n", n).Info("events successfully saved")
}
