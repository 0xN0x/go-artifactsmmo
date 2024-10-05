package artifactsmmo

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/internal/client"
	"github.com/0xN0x/go-artifactsmmo/models"
)

const (
	apiUrl = "https://api.artifactsmmo.com"
)

type ArtifactsMMO struct {
	mu       sync.Mutex
	Config   *client.ArtifactsConfig
	Token    string
	Username string
}

// NewClient creates a new client to access the ArtifactsMMO API
func NewClient(token string, username string) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:       sync.Mutex{},
		Config:   client.NewConfig(&http.Client{}, token, username),
		Token:    token,
		Username: username,
	}
}

// NewClientWithCustomHttpClient creates a new client with a custom http.Client, mainly used for testing
func NewClientWithCustomHttpClient(token string, username string, httpClient *http.Client) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:       sync.Mutex{},
		Config:   client.NewConfig(httpClient, token, username),
		Token:    token,
		Username: username,
	}
}

// Start a fight against a monster on the character's map.
func (c *ArtifactsMMO) Fight() (*models.CharacterFight, error) {
	var fight models.CharacterFight

	_, err := api.NewRequest(c.Config, &fight, "POST", fmt.Sprintf("%s/my/%s/action/fight", apiUrl, c.Username), nil).Run()
	if err != nil {
		return nil, err
	}

	return &fight, nil
}

// Retrieve the details of a character.
func (c *ArtifactsMMO) GetCharacterInfo() (*models.Character, error) {
	var character models.Character

	_, err := api.NewRequest(c.Config, &character, "GET", fmt.Sprintf("%s/characters/%s", apiUrl, c.Username), nil).Run()
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// Moves a character on the map using the map's X and Y position.
func (c *ArtifactsMMO) Move(x int, y int) (*models.CharacterMovementData, error) {
	var move models.CharacterMovementData

	body := models.Movement{X: x, Y: y}
	res, err := api.NewRequest(c.Config, &move, "POST", fmt.Sprintf("%s/my/%s/action/move", apiUrl, c.Username), body).Run()

	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrMapNotFound
	case 490:
		return nil, models.ErrAlreadyAtDestination
	}

	return &move, nil
}

// Equip an item on your character.
func (c *ArtifactsMMO) Equip(code string, slot models.Slot, quantity int) (*models.EquipRequest, error) {
	var equip models.EquipRequest

	body := models.ItemInventory{Code: code, Slot: slot, Quantity: quantity}
	res, err := api.NewRequest(c.Config, &equip, "POST", fmt.Sprintf("%s/my/%s/action/equip", apiUrl, c.Username), body).Run()

	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 484:
		return nil, models.ErrTooMuchConsumables
	case 485:
		return nil, models.ErrItemAlreadyEquiped
	case 491:
		return nil, models.ErrSlotNotEmpty
	case 496:
		return nil, models.ErrLevelTooLow
	}

	return &equip, nil
}

// Unequip an item on your character.
func (c *ArtifactsMMO) Unequip(slot models.Slot, quantity int) (*models.EquipRequest, error) {
	var unequip models.EquipRequest

	body := models.RemoveItemInventory{Slot: slot, Quantity: quantity}
	res, err := api.NewRequest(c.Config, &unequip, "POST", fmt.Sprintf("%s/my/%s/action/unequip", apiUrl, c.Username), body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	}

	return &unequip, nil
}

// Harvest a resource on the character's map.
func (c *ArtifactsMMO) Gather() (*models.SkillData, error) {
	var skill models.SkillData

	res, err := api.NewRequest(c.Config, &skill, "POST", fmt.Sprintf("%s/my/%s/action/gathering", apiUrl, c.Username), nil).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 493:
		return nil, models.ErrInsufficientSkillLevel
	case 598:
		return nil, models.ErrRessourceNotFound
	}

	return &skill, nil
}
