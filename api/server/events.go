package server

import (
	"encoding/json"
)

type EncodedEvent struct {
	Timestamp   uint64          `json:"timestmap"`
	TrackingID  string          `json:"id"`
	Type        string          `json:"type"`
	AnonymousID string          `json:"aid"`
	Data        json.RawMessage `json:"data"`
}

type Event struct {
	Timestamp   uint64      `json:"timestmap"`
	TrackingID  string      `json:"id"`
	Type        string      `json:"type"`
	AnonymousID string      `json:"aid"`
	Data        interface{} `json:"data"`
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
