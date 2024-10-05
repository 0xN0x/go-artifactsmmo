package models

type Fight struct {
	// Main
	Xp     int    `json:"xp"`
	Gold   int    `json:"gold"`
	Turns  int    `json:"turns"`
	Result string `json:"result"`

	// Drops
	Drops []SimpleItem `json:"drops"`

	// Monster hits
	MonsterBlockedHits struct {
		Fire  int `json:"fire"`
		Earth int `json:"earth"`
		Water int `json:"water"`
		Air   int `json:"air"`
		Total int `json:"total"`
	} `json:"monster_blocked_hits"`

	// Player hits
	PlayerBlockedHits struct {
		Fire  int `json:"fire"`
		Earth int `json:"earth"`
		Water int `json:"water"`
		Air   int `json:"air"`
		Total int `json:"total"`
	} `json:"player_blocked_hits"`

	Logs []string `json:"logs"`
}

type CharacterFight struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Fight     Fight     `json:"fight"`
	Character Character `json:"character"`
}
