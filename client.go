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
	mu     sync.Mutex
	Config *client.ArtifactsConfig
}

// NewClient creates a new client to access the ArtifactsMMO API
func NewClient(token string, username string) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:     sync.Mutex{},
		Config: client.NewConfig(&http.Client{}, token, username),
	}
}

// NewClientWithCustomHttpClient creates a new client with a custom http.Client, mainly used for testing
func NewClientWithCustomHttpClient(token string, username string, httpClient *http.Client) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:     sync.Mutex{},
		Config: client.NewConfig(httpClient, token, username),
	}
}

// Start a fight against a monster on the character's map.
func (c *ArtifactsMMO) Fight() (*models.CharacterFight, error) {
	var fight models.CharacterFight

	_, err := api.NewRequest(c.Config, &fight, "POST", fmt.Sprintf("%s/my/%s/action/fight", apiUrl, c.Config.GetUsername()), nil).Run()
	if err != nil {
		return nil, err
	}

	return &fight, nil
}

// Retrieve the details of a character.
func (c *ArtifactsMMO) GetCharacterInfo() (*models.Character, error) {
	var character models.Character

	_, err := api.NewRequest(c.Config, &character, "GET", fmt.Sprintf("%s/characters/%s", apiUrl, c.Config.GetUsername()), nil).Run()
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// Moves a character on the map using the map's X and Y position.
func (c *ArtifactsMMO) Move(x int, y int) (*models.CharacterMovementData, error) {
	var move models.CharacterMovementData

	body := models.Movement{X: x, Y: y}
	res, err := api.NewRequest(c.Config, &move, "POST", fmt.Sprintf("%s/my/%s/action/move", apiUrl, c.Config.GetUsername()), body).Run()

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
	res, err := api.NewRequest(c.Config, &equip, "POST", fmt.Sprintf("%s/my/%s/action/equip", apiUrl, c.Config.GetUsername()), body).Run()

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
	res, err := api.NewRequest(c.Config, &unequip, "POST", fmt.Sprintf("%s/my/%s/action/unequip", apiUrl, c.Config.GetUsername()), body).Run()
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

	res, err := api.NewRequest(c.Config, &skill, "POST", fmt.Sprintf("%s/my/%s/action/gathering", apiUrl, c.Config.GetUsername()), nil).Run()
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

// Accepting a new task.
func (c *ArtifactsMMO) AcceptNewTask() (*models.TaskData, error) {
	var task models.TaskData

	res, err := api.NewRequest(c.Config, &task, "POST", fmt.Sprintf("%s/my/%s/action/task/new", apiUrl, c.Config.GetUsername()), nil).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &task, nil
}

// Complete a task.
func (c *ArtifactsMMO) CompleteTask() (*models.TaskRewardData, error) {
	var taskReward models.TaskRewardData

	res, err := api.NewRequest(c.Config, &taskReward, "POST", fmt.Sprintf("%s/my/%s/action/task/complete", apiUrl, c.Config.GetUsername()), nil).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &taskReward, nil
}

// Exchange 6 tasks coins for a random reward. Rewards are exclusive items or resources.
func (c *ArtifactsMMO) TaskExchange() (*models.TaskRewardData, error) {
	var task models.TaskRewardData

	res, err := api.NewRequest(c.Config, &task, "POST", fmt.Sprintf("%s/my/%s/action/task/exchange", apiUrl, c.Config.GetUsername()), nil).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &task, nil
}

// Trading items with a Tasks Master.
func (c *ArtifactsMMO) TaskTrade(code string, quantity int) (*models.TaskTradeData, error) {
	var task models.TaskTradeData

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config, &task, "POST", fmt.Sprintf("%s/my/%s/action/task/trade", apiUrl, c.Config.GetUsername()), body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &task, nil
}

// Cancel a task for 1 tasks coin.
func (c *ArtifactsMMO) TaskCancel() (*models.TaskCancelled, error) {
	var task models.TaskCancelled

	res, err := api.NewRequest(c.Config, &task, "POST", fmt.Sprintf("%s/my/%s/action/task/cancel", apiUrl, c.Config.GetUsername()), nil).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 598:
		return nil, models.ErrTaskmasterNotFound
	}

	return &task, nil
}

func (c *ArtifactsMMO) Craft(code string, quantity int) (*models.SkillData, error) {
	var ret models.SkillData

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config, &ret, "POST", fmt.Sprintf("%s/my/%s/action/crafting", apiUrl, c.Config.GetUsername()), body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrCraftNotFound
	case 493:
		return nil, models.ErrInsufficientSkillLevel
	case 598:
		return nil, models.ErrWorkshopNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) DepositBank(code string, quantity int) (*models.BankItemTransaction, error) {
	var ret models.BankItemTransaction

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config, &ret, "POST", fmt.Sprintf("%s/my/%s/action/bank/deposit", apiUrl, c.Config.GetUsername()), body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 461:
		return nil, models.ErrTransactionInProgress
	case 462:
		return nil, models.ErrBankFull
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) DepositBankGold(quantity int) (*models.BankGoldTransaction, error) {
	var ret models.BankGoldTransaction

	body := models.Gold{Quantity: quantity}
	res, err := api.NewRequest(c.Config, &ret, "POST", fmt.Sprintf("%s/my/%s/action/bank/deposit/gold", apiUrl, c.Config.GetUsername()), body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 461:
		return nil, models.ErrTransactionInProgress
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) WithdrawBank(code string, quantity int) (*models.BankItemTransaction, error) {
	var ret models.BankItemTransaction

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config, &ret, "POST", fmt.Sprintf("%s/my/%s/action/bank/withdraw", apiUrl, c.Config.GetUsername()), body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 461:
		return nil, models.ErrTransactionInProgress
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) WithdrawBankGold(quantity int) (*models.BankGoldTransaction, error) {
	var ret models.BankGoldTransaction

	body := models.Gold{Quantity: quantity}
	res, err := api.NewRequest(c.Config, &ret, "POST", fmt.Sprintf("%s/my/%s/action/bank/withdraw/gold", apiUrl, c.Config.GetUsername()), body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 460:
		return nil, models.ErrInsufficientGold
	case 461:
		return nil, models.ErrTransactionInProgress
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}
