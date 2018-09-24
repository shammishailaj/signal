package db

import (
	"encoding/json"

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
