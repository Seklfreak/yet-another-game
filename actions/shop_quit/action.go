package shop_quit

import (
	"fmt"

	"github.com/Seklfreak/yet-another-game/color"
	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Move on"
}

func (a *Action) Do(state *models.State) bool {
	fmt.Printf("%s\"Good luck on your adventures!\"%s\n", color.Purple, color.Reset)
	return true
}
