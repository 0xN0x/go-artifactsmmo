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
		Config: client.NewConfig(&http.Client{}, apiUrl, token, username),
	}
}

// NewClientWithCustomHttpClient creates a new client with a custom http.Client, mainly used for testing
func NewClientWithCustomHttpClient(token string, username string, httpClient *http.Client) *ArtifactsMMO {
	return &ArtifactsMMO{
		mu:     sync.Mutex{},
		Config: client.NewConfig(httpClient, apiUrl, token, username),
	}
}

// Start a fight against a monster on the character's map.
func (c *ArtifactsMMO) Fight() (*models.CharacterFight, error) {
	var fight models.CharacterFight

	_, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/fight", c.Config.GetUsername())).SetResultStruct(&fight).Run()
	if err != nil {
		return nil, err
	}

	return &fight, nil
}

// Retrieve the details of a character.
func (c *ArtifactsMMO) GetCharacterInfo(name string) (*models.Character, error) {
	var character models.Character

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/characters/%s", name)).SetResultStruct(&character).Run()
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// Moves a character on the map using the map's X and Y position.
func (c *ArtifactsMMO) Move(x int, y int) (*models.CharacterMovementData, error) {
	var move models.CharacterMovementData

	body := models.Movement{X: x, Y: y}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/move", c.Config.GetUsername())).SetResultStruct(&move).SetBody(body).Run()
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
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/equip", c.Config.GetUsername())).SetResultStruct(&equip).SetBody(body).Run()
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
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/unequip", c.Config.GetUsername())).SetResultStruct(&unequip).SetBody(body).Run()
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

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/gathering", c.Config.GetUsername())).SetResultStruct(&skill).Run()
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

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/new", c.Config.GetUsername())).SetResultStruct(&task).Run()
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

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/complete", c.Config.GetUsername())).SetResultStruct(&taskReward).Run()
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

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/exchange", c.Config.GetUsername())).SetResultStruct(&task).Run()
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
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/trade", c.Config.GetUsername())).SetResultStruct(&task).SetBody(body).Run()
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

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/task/cancel", c.Config.GetUsername())).SetResultStruct(&task).Run()
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
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/crafting", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
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

func (c *ArtifactsMMO) Recycling(code string, quantity int) (*models.Recycling, error) {
	var ret models.Recycling

	body := models.SimpleItem{Code: code, Quantity: quantity}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/recycling", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 404:
		return nil, models.ErrItemNotFound
	case 473:
		return nil, models.ErrCantBeRecycled
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

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/deposit", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
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
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/deposit/gold", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
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
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/withdraw", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
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
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/withdraw/gold", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
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

func (c *ArtifactsMMO) BuyBankExpansion() (*models.BankTransaction, error) {
	var ret models.BankTransaction

	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/bank/buy_expansion", c.Config.GetUsername())).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrBankNotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) BuyGE(code string, quantity int, price int) (*models.GETransaction, error) {
	var ret models.GETransaction

	body := models.GEItem{Code: code, Quantity: quantity, Price: price}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/ge/buy", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 479:
		return nil, models.ErrTooManyItems
	case 480:
		return nil, models.ErrNoStock
	case 482:
		return nil, models.ErrNoItem
	case 483:
		return nil, models.ErrTransactionInProgress
	case 486:
		return nil, models.ErrTransactionCharacter
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrGENotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) SellGE(code string, quantity int, price int) (*models.GETransaction, error) {
	var ret models.GETransaction

	body := models.GEItem{Code: code, Quantity: quantity, Price: price}
	res, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/ge/sell", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 479:
		return nil, models.ErrTooManyItems
	case 480:
		return nil, models.ErrNoStock
	case 482:
		return nil, models.ErrNoItem
	case 483:
		return nil, models.ErrTransactionInProgress
	case 486:
		return nil, models.ErrTransactionCharacter
	case 492:
		return nil, models.ErrInsufficientGold
	case 598:
		return nil, models.ErrGENotFound
	}

	return &ret, nil
}

func (c *ArtifactsMMO) DeleteItem(code string, quantity int) (*models.ItemReponse, error) {
	var ret models.ItemReponse

	body := models.SimpleItem{Code: code, Quantity: quantity}
	_, err := api.NewRequest(c.Config).SetMethod("POST").SetURL(fmt.Sprintf("/my/%s/action/delete", c.Config.GetUsername())).SetResultStruct(&ret).SetBody(body).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// Retrieve the details of all the characters.
func (c *ArtifactsMMO) GetMyCharactersInfo() (*[]models.Character, error) {
	var ret []models.Character

	_, err := api.NewRequest(c.Config).SetMethod("GET").SetURL("/my/characters").SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// Retrieve the details of an achievement
func (c *ArtifactsMMO) GetAchievement(code string) (*models.BaseAchievement, error) {
	var ret models.BaseAchievement

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/achievements/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrAchievementNotFound
	}

	return &ret, nil
}

// Retrieve the details of a map tile
func (c *ArtifactsMMO) GetMap(x int, y int) (*models.MapSchema, error) {
	var ret models.MapSchema

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/maps/%d/%d", x, y)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrMapNotFound
	}

	return &ret, nil
}

// Fetch maps details.
func (c *ArtifactsMMO) GetMaps(contentCode, contentType string, page, size int) (any, error) {
	var ret any

	_, err := api.NewRequest(c.Config).
		SetMethod("GET").
		SetURL("/maps").
		SetResultStruct(&ret).
		SetParam("content_code", contentCode).
		SetParam("content_type", contentType).
		SetParam("page", string(page)).
		SetParam("size", string(size)).Run()
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// Retrieve the details of an item
func (c *ArtifactsMMO) GetItem(code string) (*models.SingleItem, error) {
	var ret models.SingleItem

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/items/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrItemNotFound
	}

	return &ret, nil
}

// Retrieve the details of a monster
func (c *ArtifactsMMO) GetMonster(code string) (*models.Monster, error) {
	var ret models.Monster

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/monsters/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrMonsterNotFound
	}

	return &ret, nil
}

// Retrieve the details of a ge item
func (c *ArtifactsMMO) GetGEItem(code string) (*models.GEItems, error) {
	var ret models.GEItems

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/ge/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrItemNotFound
	}

	return &ret, nil
}

// Retrieve the details of a task
func (c *ArtifactsMMO) GetTask(code string) (*models.TaskFull, error) {
	var ret models.TaskFull

	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/tasks/list/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrTaskNotFound
	}

	return &ret, nil
}

// Retrieve the details of a task reward
func (c *ArtifactsMMO) GetTaskReward(code string) (*models.TaskRewardFull, error) {
	var ret models.TaskRewardFull

	//res, err := api.NewRequest(c.Config, &ret, "GET", fmt.Sprintf("%s/tasks/rewards/%s", apiUrl, code), nil).Run()
	res, err := api.NewRequest(c.Config).SetMethod("GET").SetURL(fmt.Sprintf("/tasks/rewards/%s", code)).SetResultStruct(&ret).Run()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, models.ErrRewardNotFound
	}

	return &ret, nil
}
