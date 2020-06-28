package gain_exp

import (
	"fmt"
	"strconv"

	"github.com/Seklfreak/yet-another-game/color"
	"github.com/Seklfreak/yet-another-game/models"
)

type Action struct {
}

func (a *Action) Key() string {
	return "Gain EXP"
}

func (a *Action) Do(state *models.State) bool {
	amount, _ := strconv.Atoi(state.ActionContext["gain_exp_amount"])

	currentLevel := state.GetLevel()

	state.Exp += amount

	newLevel := state.GetLevel()

	if newLevel > currentLevel {
		fmt.Printf("%sLevel up! You are now level %d.%s\n", color.Red, newLevel, color.Reset)
	} else {
		fmt.Printf("You gained %s%d EXP%s.\n", color.Yellow, amount, color.Reset)
	}

	return false
}
