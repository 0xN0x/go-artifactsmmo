package artifactsmmo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/0xN0x/go-artifactsmmo/models"
)

const (
	apiUrl = "https://api.artifactsmmo.com"
)

type ArtifactsMMO struct {
	mu       sync.Mutex
	token    string
	username string
	client   http.Client
}

func NewClient(token string, username string) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:       sync.Mutex{},
		token:    token,
		username: username,
		client:   http.Client{},
	}
}

func (c *ArtifactsMMO) makeRequest(method string, action string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", apiUrl, action)

	req, _ := http.NewRequest(method, url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	if method == "POST" {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := c.client.Do(req)

	switch res.StatusCode {
	case 452:
		return nil, fmt.Errorf("failed to parse token: token is missing or empty")
	case 486:
		return nil, fmt.Errorf("an action is already in progress by your character")
	case 497:
		return nil, fmt.Errorf("character inventory is full")
	case 498:
		return nil, fmt.Errorf("character not found")
	case 499:
		return nil, fmt.Errorf("character in cooldown")
	}

	return res, err
}

func (c *ArtifactsMMO) Fight() (*models.CharacterFight, error) {
	res, err := c.makeRequest("POST", fmt.Sprintf("my/%s/action/fight", c.username))

	if err != nil {
		return nil, err
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	if res.StatusCode == 598 {
		return nil, fmt.Errorf("monster not found on this map")
	}

	body, _ := io.ReadAll(res.Body)

	var fight struct {
		Data models.CharacterFight `json:"data"`
	}

	json.Unmarshal(body, &fight)

	return &fight.Data, nil
}

func (c *ArtifactsMMO) GetCharacterInfo() (*models.Character, error) {
	res, err := c.makeRequest("GET", fmt.Sprintf("characters/%s", c.username))

	if err != nil {
		return nil, err
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	body, _ := io.ReadAll(res.Body)

	var character struct {
		Data models.Character `json:"data"`
	}
	json.Unmarshal(body, &character)

	return &character.Data, nil
}
