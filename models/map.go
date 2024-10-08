package models

type MapContentType string

const (
	MapMonster       MapContentType = "monster"
	MapResource      MapContentType = "resource"
	MapWorkshop      MapContentType = "workshop"
	MapBank          MapContentType = "bank"
	MapGrandExchange MapContentType = "grand_exchange"
	MapTasksMaster   MapContentType = "tasks_master"
)

type Destination struct {
	Name    string     `json:"name"`
	Skin    string     `json:"skin"`
	X       int        `json:"x"`
	Y       int        `json:"y"`
	Content MapContent `json:"content,omitempty"`
}

type MapContent struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type Movement struct {
	X int `json:"x"`
	Y int `json:"y"`
}
