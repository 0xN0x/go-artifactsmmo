package models

type Item struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Level       int    `json:"level"`
	Type        string `json:"type"`
	SubType     string `json:"sub_type"`
	Description string `json:"description"`
	Effects     []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"effects"`

	Craft []Craft `json:"craft"`
}

type Craft struct {
	Skill    string       `json:"skill"`
	Level    int          `json:"level"`
	Quantity int          `json:"quantity"`
	Items    []SimpleItem `json:"items"`
}

type SimpleItem struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Drop struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Detail struct {
	Xp    int    `json:"xp"`
	Items []Drop `json:"items"`
}

type Gold struct {
	Quantity int `json:"quantity"`
}
