package artifactsmmo

import (
	"fmt"
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
	client   *http.Client
}

// NewClient creates a new client to access the ArtifactsMMO API
func NewClient(token string, username string) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:       sync.Mutex{},
		token:    token,
		username: username,
		client:   &http.Client{},
	}
}

// NewClientWithCustomHttpClient creates a new client with a custom http.Client, mainly used for testing
func NewClientWithCustomHttpClient(token string, username string, client *http.Client) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:       sync.Mutex{},
		token:    token,
		username: username,
		client:   client,
	}
}

// Fight starts a fight with a monster on the current map (if there is any)
func (c *ArtifactsMMO) Fight() (*models.CharacterFight, error) {
	var fight struct {
		Data models.CharacterFight `json:"data"`
	}

	_, err := NewRequest(c, &fight, "POST", fmt.Sprintf("%s/my/%s/action/fight", apiUrl, c.username), nil).Run()
	if err != nil {
		return nil, err
	}

	return &fight.Data, nil
}

// GetCharacterInfo returns the full character information
func (c *ArtifactsMMO) GetCharacterInfo() (*models.Character, error) {
	var character struct {
		Data models.Character `json:"data"`
	}

	_, err := NewRequest(c, &character, "GET", fmt.Sprintf("%s/characters/%s", apiUrl, c.username), nil).Run()
	if err != nil {
		return nil, err
	}

	return &character.Data, nil
}
