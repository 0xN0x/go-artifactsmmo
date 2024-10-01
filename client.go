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

func NewClient(token string, username string) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:       sync.Mutex{},
		token:    token,
		username: username,
		client:   &http.Client{},
	}
}

func NewClientWithCustomHttpClient(token string, username string, client *http.Client) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:       sync.Mutex{},
		token:    token,
		username: username,
		client:   client,
	}
}

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
