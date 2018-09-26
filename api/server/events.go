package server

type Event struct {
	Timestamp   int64                  `json:"timestamp"`
	TrackingID  string                 `json:"tid"`
	Type        string                 `json:"type"`
	AnonymousID string                 `json:"aid"`
	Data        map[string]interface{} `json:"data"`
}
