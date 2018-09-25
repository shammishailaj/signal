package db

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

const BcryptCost = 12

type Account struct {
	gorm.Model
	Username string
	Password string
}

type Project struct {
	gorm.Model
	Name       string
	TrackingID string
}

type AnalyticsEvent struct {
	gorm.Model
	ProjectID uint
	Project   Project

	//
	ReceivedAt  time.Time
	UserAgent   string
	IP          string
	Timestamp   time.Time       `json:"timestmap"`
	TrackingID  string          `json:"id"`
	Type        string          `json:"type"`
	AnonymousID string          `json:"aid"`
	Data        json.RawMessage `json:"data"`
}

func (ev AnalyticsEvent) GetData() (interface{}, error) {
	var err error

	switch ev.Type {
	case "page_view":
		var pv PageView
		err = json.Unmarshal(ev.Data, &pv)
		return pv, err
	default:
		return nil, errors.New(ev.Type + " unrecognized")
	}
}

type Device struct {
	Width  uint64 `json:"w"`
	Height uint64 `json:"h"`
}

type PageView struct {
	Referrer string `json:"referrer"`
	Device   Device `json:"device"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Hash     string `json:"hash"`
	Title    string `json:"title"`
	Query    string `json:"query"`
	URL      string `json:"url"`
}
