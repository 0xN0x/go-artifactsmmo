package main

import (
	"fmt"
	"os"
	"time"

	"github.com/0xN0x/go-artifactsmmo"
)

func usage() {
	fmt.Println("Usage: go run main.go <api-token> <character-name>")
}

// Make the program sleep for n seconds
func waitCooldown(seconds int) {
	time.Sleep(seconds * time.Second)
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	client := artifactsmmo.NewClient(os.Args[1], os.Args[2])
	character, err := client.GetCharacterInfo(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Welcome, %s (XP: %d/%d)!\nCurrent map: [%d,%d]\n", character.Name, character.Xp, character.MaxXp, character.X, character.Y)

	// Go in [2,0] (copper mining map)
	movement, err := client.Move(2, 0)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Moved in [%d, %d]! type=%s, code=%s\n",
		movement.Destination.X,
		movement.Destination.Y,
		movement.Destination.Content.Type,
		movement.Destination.Content.Code,
	)

	// Farm while checking inventory
	for {
		skillData, err := client.Gather()
		if err != nil {
			// If inventory is full, leave this loop to craft and sell
			if err == "character inventory is full" {
				break
			}

			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Gathered %d %s and gained %d XP!",
			skillData.Detail.Items[0].Quantity,
			skillData.Detail.Items[0].Code,
			skillData.Detail.XP,
		)

		waitCooldown(skillData.Cooldown.RemainingSeconds)
	}

	// Move to craft zone
	movement, err := client.Move(1, 5)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Moved in [%d, %d]! type=%s, code=%s\n",
		movement.Destination.X,
		movement.Destination.Y,
		movement.Destination.Content.Type,
		movement.Destination.Content.Code,
	)

	waitCooldown(movement.Cooldown.RemainingSeconds)

	// Check how much copper is in inventory
	character, err := client.GetCharacterInfo(os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	waitCooldown(character.Cooldown.RemainingSeconds)

	copperOreQuantity := 0

	for _, slot := range character.Inventory {
		if slot.Code != "copper_ore" {
			continue
		}

		copperOreQuantity = slot.Quantity
	}

	// Craft the max quantity
	craft, err := client.Craft("copper", copperOreQuantity)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Crafted %d %s!", craft.Details.Items[0].Quantity, craft.Details.Items[0].Code)

	waitCooldown(craft.Cooldown.RemainingSeconds)

	// Move to sell zone
	movement, err := client.Move(1, 5)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Moved in [%d, %d]! type=%s, code=%s\n",
		movement.Destination.X,
		movement.Destination.Y,
		movement.Destination.Content.Type,
		movement.Destination.Content.Code,
	)

	waitCooldown(movement.Cooldown.RemainingSeconds)

	// Sell all copper


	// Then store money in bank

	// Repeat

	fight, err := client.Fight()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("[%s] Fight result: +%d xp, +%d gold\n", fight.Fight.Result, fight.Fight.Xp, fight.Fight.Gold)
}
