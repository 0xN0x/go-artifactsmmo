package models

type Cooldown struct {
	TotalSeconds     int    `json:"total_seconds"`
	RemainingSeconds int    `json:"remaining_seconds"`
	StartedAt        string `json:"started_at"`
	Expiration       string `json:"expiration"`
	Reason           string `json:"reason"`
}
