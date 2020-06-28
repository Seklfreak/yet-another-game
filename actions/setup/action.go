package setup

import (
	"fmt"

	"github.com/Seklfreak/yet-another-game/color"
	"github.com/Seklfreak/yet-another-game/input"
	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Setup"
}

func (a *Action) Do(state *models.State) bool {
	if state.Name != "" {
		return false
	}

	fmt.Println("You are an explorer out to seek luck in the deeps of galactic space.")
	fmt.Println("You just got your first own ship.")

	state.Name = input.Text("Name your ship: ")

	fmt.Printf("%sLet's start the adventure! 🚀%s\n", color.Red, color.Reset)

	return false
}
