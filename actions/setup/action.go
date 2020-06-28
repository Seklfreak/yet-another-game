package setup

import (
	"fmt"

	"github.com/Seklfreak/yet-another-game/input"
	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
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

	fmt.Printf(promptui.Styler(promptui.FGRed)("Let's start the adventure! 🚀\n"))

	return false
}
