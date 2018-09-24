package main

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name string
}

type AnalyticsEvent struct {
	gorm.Model
	ProjectID int
	Project   Project

	//
	Timestamp   uint64          `json:"timestmap"`
	TrackingID  string          `json:"id"`
	Type        string          `json:"type"`
	AnonymousID string          `json:"aid"`
	Data        json.RawMessage `json:"data"`
}
