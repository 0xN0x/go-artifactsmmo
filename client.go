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

func (c *ArtifactsMMO) GetCharacterInfo() (*models.Character, error) {
	url := fmt.Sprintf("%s/characters/%s", apiUrl, c.username)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := c.client.Do(req)

	if res.StatusCode != 200 {
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
