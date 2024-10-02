package models

type Destination struct {
	Name    string     `json:"name"`
	Skin    string     `json:"skin"`
	X       int        `json:"x"`
	Y       int        `json:"y"`
	Content MapContent `json:"content, omitempty"`
}

type MapContent struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type CharacterMovementData struct {
	Cooldown    Cooldown    `json:"cooldown"`
	Destination Destination `json:"destination"`
	Character   Character   `json:"character"`
}

type Movement struct {
	X int `json:"x"`
	Y int `json:"y"`
}
