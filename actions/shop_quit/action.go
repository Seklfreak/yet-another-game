package shop_quit

import (
	"fmt"

	"github.com/Seklfreak/yet-another-game/models"
	"github.com/manifoldco/promptui"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Move on"
}

func (a *Action) Do(state *models.State) bool {
	fmt.Println(promptui.Styler(promptui.FGMagenta)("\"Good luck on your adventures!\""))
	return true
}
