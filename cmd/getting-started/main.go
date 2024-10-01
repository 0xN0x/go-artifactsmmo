package main

import (
	"fmt"
	"os"

	"github.com/0xN0x/go-artifactsmmo"
)

func usage() {
	fmt.Println("Usage: go run main.go <api-token> <character-name>")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	client := artifactsmmo.NewClient(os.Args[1], os.Args[2])
	character, err := client.GetCharacterInfo()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Welcome, %s (XP: %d/%d)!\n", character.Name, character.Xp, character.MaxXp)
}
